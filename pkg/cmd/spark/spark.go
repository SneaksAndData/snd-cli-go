// Package spark /*
package spark

import (
	"github.com/spf13/cobra"
)

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
	cmd.AddCommand(NewCmdEncrypt())
	cmd.AddCommand(NewCmdConfiguration())

	cmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")

	return cmd
}
