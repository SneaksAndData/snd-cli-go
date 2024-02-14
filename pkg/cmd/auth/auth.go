package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
)

var env, provider string

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Get internal authorization token",
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginRun()
		},
	}

	util.DisableAuthCheck(cmd)

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")

	return cmd
}

func loginRun() error {
	fmt.Println("Login")
	fmt.Printf(boxerBaseURL, env)
	panic("Not implemented")
}
