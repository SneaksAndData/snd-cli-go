package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

var value, sp string

func NewCmdEncrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Encrypt called")
		},
	}
	cmd.Flags().StringVarP(&value, "value", "v", "", "Value to encrypt")
	cmd.Flags().StringVarP(&sp, "secret-path", "s", "", "Optional Vault secret path to Spark Runtime encryption key")

	return cmd
}
