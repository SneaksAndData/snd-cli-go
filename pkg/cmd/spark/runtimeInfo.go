package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
)

var object string

func NewCmdRuntimeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime-info",
		Short: "Get the runtime info of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {

			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return runtimeInfoRun(sparkService)
		},
	}

	cmd.Flags().StringVarP(&object, "object", "o", "", "Apply a filter on the returned JSON output")

	return cmd
}

func runtimeInfoRun(sparkService *spark.Service) error {
	response, err := sparkService.GetRuntimeInfo(id)
	if err != nil {
		log.Fatalf("Failed to retrieve runtime info for run id %s: %v", id, err)
	}

	fmt.Println("Response:", response)
	return nil
}
