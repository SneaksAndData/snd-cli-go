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
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "Specify the  Job ID")

	return cmd
}
