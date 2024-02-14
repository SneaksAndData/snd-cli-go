package spark

import (
	"github.com/spf13/cobra"
)

var name string

func NewCmdConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return configurationRun()
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", " Name of the configuration to find")
	return cmd
}

func configurationRun() error {
	panic("Not implemented")
}
