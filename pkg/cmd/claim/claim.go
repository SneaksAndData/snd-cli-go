// Package claim /*
package claim

import (
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/claim/user"
)

func NewCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "claim",
		Short:   "Manage claims",
		GroupID: "claim",
	}

	cmd.AddCommand(user.NewCmdUser())
	cmd.AddCommand(NewCmdAddClaim())
	cmd.AddCommand(NewCmdGetClaim())
	cmd.AddCommand(NewCmdRemoveClaim())

	cmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")
	cmd.Flags().StringP("user", "u", "", "Specify the user ID")
	cmd.Flags().StringP("provider", "p", "", "Specify the claim provider")

	return cmd
}
