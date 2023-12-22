/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("configuration called")
		},
	}

	cmd.Flags().StringP("name", "n", "", " Name of the configuration to find")
	return cmd

}
