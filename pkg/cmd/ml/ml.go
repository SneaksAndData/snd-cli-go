package ml

import (
	"github.com/spf13/cobra"
)

func NewCmdAlgorithm() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "algorithm",
		Short:   "Manage ML algorithm jobs",
		GroupID: "ml",
	}

	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdRun())
	cmd.AddCommand(NewCmdSubmit())

	cmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")
	cmd.PersistentFlags().StringP("algorithm", "l", "", "Specify the algorithm name")

	return cmd
}
