package token

import (
	"fmt"
	"os"
	"path/filepath"
)

const folder = ".snd-cli"
const tokenFileName = "user-token.json"

type TokenCacher interface {
	CacheToken(token string) (string, error)
	ReadToken() (string, error)
}

type TokenCache struct {
	Token string
}

func (tc *TokenCache) CacheToken() (string, error) {
	filePath, err := generateTokenCachePath()
	if err != nil {
		return "", err
	}

	dirPath := filepath.Dir(filePath)
	if err := createDirectory(dirPath); err != nil {
		return "", err
	}

	if err := writeToFile(filePath, tc.Token); err != nil {
		return "", err
	}

	return tc.Token, nil
}

func (tc *TokenCache) ReadToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, folder, tokenFileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read filePath: %v", err)
	}

	return string(data), nil
}

// GenerateTokenCachePath generates the full path for the token cache file
func generateTokenCachePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %v", err)
	}
	dirPath := filepath.Join(homeDir, folder)
	return filepath.Join(dirPath, tokenFileName), nil
}

func createDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755) // Use MkdirAll to simplify
	if err != nil {
		return fmt.Errorf("failed to create token cache directory: %v", err)
	}
	return nil
}

// WriteToFile writes data to the specified file path.
func writeToFile(filePath, data string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
