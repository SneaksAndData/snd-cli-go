package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
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

	cmdutil.DisableAuthCheck(cmd)

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth_provider", "a", "azuread", "Specify the authentication provider name")

	return cmd
}

func run(args []string) error {
	// TODO: login logic, return nil if successful otherwise error
	fmt.Println(env)
	fmt.Println(provider)
	client := azure.NewClient("")
	baseURL := fmt.Sprintf("https://boxer.%s.sneaksanddata.com", env)
	t, err := boxer.GetBoxerToken(client.GetDefaultToken, provider, baseURL)
	if err != nil {
		return err
	}
	err = cmdutil.CacheToken(t)
	if err != nil {
		return err
	}
	fmt.Println(t)
	return nil
}
