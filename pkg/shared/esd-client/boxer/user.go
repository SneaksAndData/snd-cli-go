package boxer

import (
	"fmt"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/api"
)

func (c connector) AddUser(user string, provider string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)
	client := api.NewClient(token)

	return client.MakeRequest("POST", targetURL, nil)
}

func (c connector) RemoveUser(user string, provider string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/claim/%s/%s", c.claimUrl, provider, user)
	token, err := util.ReadToken()
	if err != nil {
		return "", err
	}
	client := api.NewClient(token)

	return client.MakeRequest("DELETE", targetURL, nil)
}
