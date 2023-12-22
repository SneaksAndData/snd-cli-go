package ml

import (
	"github.com/spf13/cobra"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRun()
		},
	}

	cmd.Flags().StringP("payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringP("tag", "t", "", " Client-side submission identifier")

	return cmd
}

func runRun() error {
	// TODO: add logic
	return nil
}
