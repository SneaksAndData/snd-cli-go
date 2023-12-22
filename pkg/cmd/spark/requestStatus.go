/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdRequestStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-status",
		Short: "Get the status of a Spark Job",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("requestStatus called")
		},
	}

	cmd.Flags().StringP("id", "i", "", "Beast Job ID")

	return cmd

}
