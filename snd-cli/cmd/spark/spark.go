// Package spark /*
package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

// sparkCmd represents the spark command
var sparkCmd = &cobra.Command{
	Use:   "spark",
	Short: "Manage Spark jobs",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spark called")
	},
}

func GetSparkCmd() *cobra.Command {
	return sparkCmd
}

func init() {
	sparkCmd.PersistentFlags().StringP("env", "e", "test", "Target environment")
	sparkCmd.PersistentFlags().StringP("auth_provider", "a", "azuread", "Specify the authentication provider name")
}
