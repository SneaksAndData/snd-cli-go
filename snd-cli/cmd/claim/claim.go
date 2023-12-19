/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/snd-cli/cmd/claim/user"
)

// claimCmd represents the claim command
var claimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Manage claims",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("claim called")
	},
}

func GetClaimCmd() *cobra.Command {
	return claimCmd
}

func init() {
	claimCmd.AddCommand(user.GetUserCommand())
	//cmd.RootCmd.AddCommand(claimCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// claimCmd.PersistentFlags().String("foo", "", "A help for foo")
	claimCmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	claimCmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")
	claimCmd.Flags().StringP("user", "u", "", "Specify the user ID")
	claimCmd.Flags().StringP("provider", "p", "", "Specify the claim provider")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// claimCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
