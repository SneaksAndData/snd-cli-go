package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
)

var name string

func NewCmdConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Get a deployed SparkJob configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return configurationRun(sparkService)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", " Name of the configuration to find")
	return cmd
}

func configurationRun(sparkService *spark.Service) error {
	response, err := sparkService.GetConfiguration(name)
	if err != nil {
		log.Fatalf("Failed to retrieve configuration with name %s: %v", name, err)
	}
	fmt.Println("Response:", response)
	return nil
}
