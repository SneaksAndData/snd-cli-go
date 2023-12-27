package boxer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snd-cli/pkg/cmd/util"
	"strings"
	"time"
)

type claimPayload struct {
	// Fields need to be public so that json package can see it
	Operation string            `json:"operation"`
	Claims    map[string]string `json:"claims"`
}

func (c connector) GetClaim(user string, provider string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)

	token, err := util.ReadToken()
	if err != nil {
		return "", err
	}

	return makeHTTPRequest("GET", targetURL, token, nil)
}

func (c connector) AddClaim(user string, provider string, claims []string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)

	payload, err := json.Marshal(preparePayload(claims, "Insert"))
	if err != nil {
		return "", err
	}

	token, err := util.ReadToken()
	if err != nil {
		return "", err
	}

	return makeHTTPRequest("PATCH", targetURL, token, bytes.NewBuffer(payload))
}

func (c connector) RemoveClaim(user string, provider string, claims []string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)

	payload, err := json.Marshal(preparePayload(claims, "Delete"))
	if err != nil {
		return "", err
	}

	token, err := util.ReadToken()
	if err != nil {
		return "", err
	}

	return makeHTTPRequest("PATCH", targetURL, token, bytes.NewBuffer(payload))
}

func makeHTTPRequest(method, url string, token string, payload io.Reader) (string, error) {
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

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func preparePayload(claims []string, operation string) claimPayload {
	claimsMap := make(map[string]string)
	for _, s := range claims {
		c := strings.Split(s, ":")
		claimsMap[c[0]] = c[1]

	}
	return claimPayload{
		Operation: operation,
		Claims:    claimsMap,
	}
}
