package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

var value, sp string

func NewCmdEncrypt(service Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := encryptRun(service)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&value, "value", "v", "", "Value to encrypt")
	cmd.Flags().StringVarP(&sp, "secret-path", "s", "", "Optional Vault secret path to Spark Runtime encryption key")

	return cmd
}

func encryptRun(sparkService Service) (string, error) {
	panic("Not implemented")
}
