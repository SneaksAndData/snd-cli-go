package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdRequestStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-status",
		Short: "Get the status of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return requestStatusRun(sparkService)
		},
	}

	return cmd

}

func requestStatusRun(sparkService *spark.Service) error {
	response, err := sparkService.GetLifecycleStage(id)
	if err != nil {
		log.Fatalf("Failed to retrieve lifecycle stage for run id %s: %v", id, err)
	}

	fmt.Println("Response:", response)
	return nil
}
