/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runtimeInfoCmd represents the runtimeInfo command
var runtimeInfoCmd = &cobra.Command{
	Use:   "runtime-info",
	Short: "Get the runtime info of a Spark Job",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runtimeInfo called")
	},
}

func init() {
	sparkCmd.AddCommand(runtimeInfoCmd)
	runtimeInfoCmd.Flags().StringP("id", "i", "", "Beast Job ID")
	runtimeInfoCmd.Flags().StringP("object", "o", "", "Apply a filter on the returned JSON output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runtimeInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runtimeInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
