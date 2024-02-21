package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
)

var name string

func NewCmdConfiguration(factory ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := factory(env)
			if err != nil {
				return err
			}
			resp, err := configurationRun(service, name)
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

		return spark.SubmissionConfiguration{}, fmt.Errorf("failed to retrieve configuration with name %s: %v", name, err)
	}
	return response, nil
}
