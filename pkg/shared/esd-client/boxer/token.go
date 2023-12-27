package boxer

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type ExternalToken struct {
	GetToken func() (string, error)
	Provider string
	Retry    bool
}

func (c connector) GetToken() (string, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	token, err := c.auth.GetToken()
	if err != nil {
		return "", err
	}
	bearer := "Bearer " + token
	targetURL := fmt.Sprintf("%s/token/%s", c.tokenUrl, c.auth.Provider)
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}
	addHeaders(req, bearer)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := resp.Body.Close(); err != nil {
		return "", err
	}

	return string(body), nil
}

func addHeaders(r *http.Request, token string) {
	r.Header.Add("Authorization", token)
	r.Header.Add("Content-Type", "application/json")
}
