// Package claim /*
package claim

import (
	"github.com/spf13/cobra"
)

var env, authProvider, userId, claimProvider string

func NewCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "claim",
		Short:   "Manage claims",
		GroupID: "claim",
	}

	cmd.AddCommand(NewCmdUser())
	cmd.AddCommand(NewCmdAddClaim())
	cmd.AddCommand(NewCmdGetClaim())
	cmd.AddCommand(NewCmdRemoveClaim())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&userId, "user", "u", "", "Specify the user ID")
	cmd.PersistentFlags().StringVarP(&claimProvider, "claims-provider", "p", "", "Specify the claim provider")

	return cmd
}
