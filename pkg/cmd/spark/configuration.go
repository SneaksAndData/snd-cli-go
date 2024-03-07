package spark

import (
	"encoding/json"
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var name string

func NewCmdConfiguration(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := configurationRun(service.(*spark.Service), name)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", " Name of the configuration to find")
	return cmd
}

func configurationRun(sparkService Service, name string) (string, error) {
	response, err := sparkService.GetConfiguration(name)
	if err != nil {

		return "", fmt.Errorf("failed to retrieve configuration with name %s: %w", name, err)
	}
	m, err := json.Marshal(&response)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}
	return string(m), nil
}
