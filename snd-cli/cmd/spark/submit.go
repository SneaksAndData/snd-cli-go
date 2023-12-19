/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Runs the provided Beast V3 job with optional overrides",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("submit called")
	},
}

func init() {
	sparkCmd.AddCommand(submitCmd)
	submitCmd.Flags().StringP("job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	submitCmd.Flags().StringP("overrides", "o", "", "Overrides for the provided job name")
	submitCmd.Flags().StringP("client-tag", "t", "", "Client tag for this submission")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
