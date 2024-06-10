package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var value, sp string

func NewCmdEncrypt(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value from a file or stdin using encryption key from a corresponding Spark Runtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := encryptRun(service.(*spark.Service))
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
