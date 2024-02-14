package ml

import (
	"github.com/spf13/cobra"
)

var env, authProvider, algorithm string

func NewCmdAlgorithm() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "algorithm",
		Short:   "Manage ML algorithm jobs",
		GroupID: "ml",
	}

	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdRun())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "l", "", "Specify the algorithm name")

	return cmd
}
