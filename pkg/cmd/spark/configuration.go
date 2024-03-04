package spark

import (
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
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, authService)
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

func configurationRun(sparkService Service, name string) (spark.SubmissionConfiguration, error) {
	response, err := sparkService.GetConfiguration(name)
	if err != nil {

		return spark.SubmissionConfiguration{}, fmt.Errorf("failed to retrieve configuration with name %s: %w", name, err)
	}
	return response, nil
}
