package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/azure"
	"snd-cli/pkg/shared/esd-client/boxer"
)

var env, provider string

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Get internal authorization token",
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args)
		},
	}

	util.DisableAuthCheck(cmd)

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth_provider", "a", "azuread", "Specify the authentication provider name")

	return cmd
}

func run(args []string) error {
	// TODO: login logic, return nil if successful otherwise error
	// Create an ExternalToken instance
	azureToken := boxer.ExternalToken{
		GetToken: azure.GetDefaultToken,
		Provider: "azuread",
		Retry:    true,
	}
	url := fmt.Sprintf("https://boxer.%s.sneaksanddata.com", env)

	input := boxer.Input{
		TokenUrl: url,
		ClaimUrl: "",
		Auth:     azureToken,
	}

	var boxerConn boxer.Token

	boxerConn = boxer.NewConnector(input)
	token, err := boxerConn.GetToken()
	if err != nil {
		return err
	}
	fmt.Println(token)
	err = util.CacheToken(token)
	if err != nil {
		return err
	}
	return nil
}
