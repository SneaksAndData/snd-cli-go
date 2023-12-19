// Package ml /*
package ml

import (
	"fmt"
	"github.com/spf13/cobra"
)

// mlCmd represents the ml command
var mlCmd = &cobra.Command{
	Use:   "algorithm",
	Short: "Manage ML algorithm jobs",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ml called")
	},
}

func GetMlCmd() *cobra.Command {
	return mlCmd
}

func init() {
	mlCmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	mlCmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")
	mlCmd.PersistentFlags().StringP("algorithm", "l", "", "Specify the algorithm name")
}
