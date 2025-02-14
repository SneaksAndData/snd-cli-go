package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"snd-cli/pkg/cmd/util/file"
	"time"
)

const folder = ".snd-cli"
const tokenFileName = "user-token.json"

// AuthService defines the interface for an authentication service.
// It requires an implementation of GetBoxerToken method that, when called,
// returns a token as a string and any error encountered.
type AuthService interface {
	GetBoxerToken() (string, error)
}

// Provider manages token operations, including caching and retrieval,
// utilizing an AuthService to obtain tokens when necessary.
type Provider struct {
	token       string      // token holds the most recent authentication token obtained.
	ttl         time.Time   // ttl represents the time-to-live for the current token.
	env         string      // env represents the environment for which the token was obtained.
	authService AuthService // authService is an instance of AuthService used to obtain authentication tokens when required.
	cachePath   file.File   // path to the file where the token will be cached
}

// tokenCache struct is used for storing a token and its expiry time
// in a cache (such as a file).
type tokenCache struct {
	Token string    `json:"token"`
	TTL   time.Time `json:"ttl"`
	Env   string    `json:"env"`
}

// NewProvider creates a new instance of Provider using the provided AuthService.
// The AuthService is used to obtain authentication tokens when they are not available in the cache or have expired.
// The env parameter is used to ensure that the token is valid for the correct environment.
func NewProvider(authService AuthService, env string) (*Provider, error) {
	filePath, err := file.GenerateFilePathWithBaseHome(folder, tokenFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate the path for token cache: %w", err)
	}
	return &Provider{
		authService: authService,
		cachePath:   file.File{FilePath: filePath},
		env:         env,
	}, nil
}

// saveTokenToCache serializes the current token and TTL into a JSON format and writes this data to a cache file.
func (p *Provider) saveTokenToCache() error {
	t := tokenCache{
		Token: p.token,
		TTL:   p.ttl,
		Env:   p.env,
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return p.cachePath.WriteToFile(string(data))
}

// getTokenFromCache attempts to read the token and its TTL from a cache file.
// If the token is still valid (not expired), it sets the token and TTL fields on the Provider and returns nil.
// If the token is expired or if there is any issue reading from the cache, it returns an error.
func (p *Provider) getTokenFromCache() error {
	filePath, err := file.GenerateFilePathWithBaseHome(folder, tokenFileName)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var cache tokenCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return err // Invalid cache, possibly corrupted.
	}

	// Check if the token in cache is still valid and for the correct environment.
	if time.Now().Before(cache.TTL) && cache.Env == p.env {
		p.token = cache.Token
		p.ttl = cache.TTL
		return nil
	}

	return errors.New("token expired or environment mismatch")
}

// GetToken checks the current token's validity and returns it if it's still valid.
// If the token is expired or not present, it attempts to load a valid token from the cache.
// If no valid token is available in the cache, it fetches a new token using the authService,
// updates the token and its TTL, caches the new token, and returns it.
func (p *Provider) GetToken() (string, error) {
	if p.token == "" || time.Now().After(p.ttl) {
		// log.Println("Reading token from cache")
		if err := p.getTokenFromCache(); err == nil {
			return p.token, nil
		}
	}
	// Either cache is empty, or token is expired, fetch a new one.
	// log.Println("Cached token not existent or expired, retrieving new token")
	token, err := p.authService.GetBoxerToken()
	if err != nil {
		return "", err
	}
	p.token = token
	p.ttl = time.Now().Add(time.Hour) // Assuming TTL is 1 hour.
	if err := p.saveTokenToCache(); err != nil {
		return token, err
	}

	return p.token, nil
}

// GetUserFromToken extracts the user information from the JWT token stored in the Provider.
// It parses the token, extracts the claims, and retrieves the user claim.
// If the token is invalid or the user claim is not found, it logs the error and returns an empty string.
func (p *Provider) GetUserFromToken() string {
	user := ""

	token, err := p.GetToken()
	if err != nil {
		fmt.Println("failed retrieve token: %w", err)
		return user
	}

	t, _ := jwt.Parse(token, nil)

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		if userClaim, ok := claims["boxer.sneaksanddata.com/user"].(string); ok {
			user = userClaim
		} else {
			fmt.Println("User claim not found or not a string")
		}
	} else {
		fmt.Println("Invalid token claims")
	}
	return user
}
