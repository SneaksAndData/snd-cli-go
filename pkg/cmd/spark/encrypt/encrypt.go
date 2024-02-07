/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package encrypt

import (
	"github.com/spf13/cobra"
)

var value, sp, env, authProvider string

func NewCmdEncrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			return encryptRun()
		},
		GroupID: "spark",
	}
	cmd.Flags().StringVarP(&value, "value", "v", "", "Value to encrypt")
	cmd.Flags().StringVarP(&sp, "secret-path", "s", "", "Optional Vault secret path to Spark Runtime encryption key")
	cmd.Flags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.Flags().StringVarP(&authProvider, "auth_provider", "a", "azuread", "Specify the authentication provider name")

	return cmd
}

func encryptRun() error {
	return nil
}
