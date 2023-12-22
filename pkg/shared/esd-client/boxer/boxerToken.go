// Package boxer Connector for Boxer Auth API.
package boxer

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetBoxerToken(getToken func() (string, error), provider string, baseUrl string) (string, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	azureAdToken, err := getToken()
	if err != nil {
		return "", err
	}
	bearer := "Bearer " + azureAdToken
	targetURL := fmt.Sprintf("%s/token/%s", baseUrl, provider)

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", bearer)
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
