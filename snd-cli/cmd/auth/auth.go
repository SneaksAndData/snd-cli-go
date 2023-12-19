// Package auth /*
package auth

import (
	"fmt"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "login",
	Short: "Get internal authorization token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func GetAuthCmd() *cobra.Command {
	return authCmd
}

func init() {
	//cmd.RootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	authCmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	authCmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
