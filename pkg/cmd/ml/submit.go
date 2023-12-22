package ml

import (
	"github.com/spf13/cobra"
)

func NewCmdSubmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit a result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return submitRun()
		},
	}
	cmd.Flags().StringP("cause", "c", "", "Cause for submitting the result")
	cmd.Flags().StringP("message", "m", "", "Result message")
	cmd.Flags().StringP("sas_uri", "u", "", "Sas Uri")

	return cmd
}

func submitRun() error {
	// TODO: add logic
	return nil
}
