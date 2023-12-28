package boxer

import (
	"fmt"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/http"
)

func (c connector) AddUser(user string, provider string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)
	return http.MakeRequest("POST", targetURL, token, nil)
}

func (c connector) RemoveUser(user string, provider string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)
	token, err := util.ReadToken()
	if err != nil {
		return "", err
	}
	return http.MakeRequest("DELETE", targetURL, token, nil)
}
