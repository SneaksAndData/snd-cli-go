/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

var object string

func NewCmdRuntimeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime-info",
		Short: "Get the runtime info of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runtimeInfoRun()
		},
	}

	cmd.Flags().StringVarP(&object, "object", "o", "", "Apply a filter on the returned JSON output")

	return cmd
}

func runtimeInfoRun() error {
	fmt.Println("Runtime info")

	return nil

}
