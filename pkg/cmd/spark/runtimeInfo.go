/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdRuntimeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime-info",
		Short: "Get the runtime info of a Spark Job",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("runtimeInfo called")
		},
	}

	cmd.Flags().StringP("id", "i", "", "Beast Job ID")
	cmd.Flags().StringP("object", "o", "", "Apply a filter on the returned JSON output")

	return cmd
}
