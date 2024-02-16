package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
)

var value, sp string

func NewCmdEncrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return encryptRun(sparkService)
		},
	}
	cmd.Flags().StringVarP(&value, "value", "v", "", "Value to encrypt")
	cmd.Flags().StringVarP(&sp, "secret-path", "s", "", "Optional Vault secret path to Spark Runtime encryption key")

	return cmd
}

func encryptRun(sparkService *spark.Service) error {
	panic("Not implemented")
}
