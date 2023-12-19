/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

// requestStatusCmd represents the requestStatus command
var requestStatusCmd = &cobra.Command{
	Use:   "request-status",
	Short: "Get the status of a Spark Job",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("requestStatus called")
	},
}

func init() {
	sparkCmd.AddCommand(requestStatusCmd)
	requestStatusCmd.Flags().StringP("id", "i", "", "Beast Job ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
