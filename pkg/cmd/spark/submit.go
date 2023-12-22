/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"github.com/spf13/cobra"
)

func NewCmdSubmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Runs the provided Beast V3 job with optional overrides",
	}

	cmd.Flags().StringP("job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	cmd.Flags().StringP("overrides", "o", "", "Overrides for the provided job name")
	cmd.Flags().StringP("client-tag", "t", "", "Client tag for this submission")

	return cmd
}
