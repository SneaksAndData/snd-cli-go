/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package ml

import (
	"fmt"

	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a result for a ML Algorithm run",
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
	mlCmd.AddCommand(submitCmd)
	getCmd.Flags().StringP("cause", "c", "", "Cause for submitting the result")
	getCmd.Flags().StringP("message", "m", "", "Result message")
	getCmd.Flags().StringP("sas_uri", "u", "", "Sas Uri")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
