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
			return run()
		},
	}

	util.DisableAuthCheck(cmd)

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth_provider", "a", "azuread", "Specify the authentication provider name")

	return cmd
}

func run() error {
	fmt.Println("Login")
	fmt.Printf(boxerBaseURL, env)
	return nil
}
