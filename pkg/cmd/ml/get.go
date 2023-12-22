package ml

import (
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRun()
		},
	}
	cmd.Flags().StringP("id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func getRun() error {
	// TODO: add logic
	return nil
}
