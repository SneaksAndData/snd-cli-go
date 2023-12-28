package http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeRequest(method, url string, token string, payload io.Reader) (string, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	bearer := "Bearer " + token
	addHeaders(req, bearer)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("authorization failed, please run 'snd login' to refresh your token")
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func addHeaders(r *http.Request, token string) {
	r.Header.Add("Authorization", token)
	r.Header.Add("Content-Type", "application/json")
}
