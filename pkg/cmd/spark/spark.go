package spark

import (
	"github.com/spf13/cobra"
)

var env, authProvider, id string

func NewCmdSpark() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spark",
		Short:   "Manage Spark jobs",
		GroupID: "spark",
	}

	cmd.AddCommand(NewCmdSubmit())
	cmd.AddCommand(NewCmdRuntimeInfo())
	cmd.AddCommand(NewCmdRequestStatus())
	cmd.AddCommand(NewCmdLogs())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth_provider", "a", "azuread", "Specify the authentication provider name")
	cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")

	return cmd
}
