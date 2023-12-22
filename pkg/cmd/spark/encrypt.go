/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package spark

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdEncrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("encrypt called")
		},
	}
	cmd.Flags().StringP("value", "v", "", "Value to encrypt")
	cmd.Flags().StringP("secret-path", "s", "", "Optional Vault secret path to Spark Runtime encryption key")

	return cmd
}
