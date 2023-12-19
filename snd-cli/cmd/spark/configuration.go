/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configurationCmd represents the configuration command
var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Get a deployed SparkJob configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configuration called")
	},
}

func init() {
	sparkCmd.AddCommand(configurationCmd)
	configurationCmd.Flags().StringP("name", "n", "", " Name of the configuration to find")
}
