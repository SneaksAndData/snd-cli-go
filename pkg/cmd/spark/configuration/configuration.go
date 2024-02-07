/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package configuration

import (
	"github.com/spf13/cobra"
)

var name, env, authProvider string

func NewCmdConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return configurationRun()
		},
		GroupID: "spark",
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", " Name of the configuration to find")
	cmd.Flags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.Flags().StringVarP(&authProvider, "auth_provider", "a", "azuread", "Specify the authentication provider name")
	return cmd
}

func configurationRun() error {
	return nil
}
