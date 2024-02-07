/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"github.com/spf13/cobra"
)

var jobName, clientTag string
var overrides string

func NewCmdSubmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Runs the provided Beast V3 job with optional overrides",
	}

	cmd.Flags().StringVarP(&jobName, "job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	cmd.Flags().StringVarP(&overrides, "overrides", "o", "", "Overrides for the provided job name")
	cmd.Flags().StringVarP(&clientTag, "client-tag", "t", "", "Client tag for this submission")

	return cmd
}
