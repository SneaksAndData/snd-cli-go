package boxer

import (
	"fmt"
	"snd-cli/pkg/shared/esd-client/http"
)

type ExternalToken struct {
	GetToken func() (string, error)
	Provider string
	Retry    bool
}

func (c connector) GetToken() (string, error) {
	targetURL := fmt.Sprintf("%s/token/%s", c.tokenUrl, c.auth.Provider)

	token, err := c.auth.GetToken()
	if err != nil {
		return "", err
	}

	return http.MakeRequest("GET", targetURL, token, nil)
}
