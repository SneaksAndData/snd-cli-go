package ml

import (
	"fmt"
	algorithms "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"log"
)

var id string

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {

			var algorithmService, err = InitAlgorithmService(fmt.Sprintf(crystalBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return getRun(algorithmService)
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func getRun(algorithmService *algorithms.Service) error {
	response, err := algorithmService.RetrieveRun(id, algorithm)
	if err != nil {
		log.Fatalf("Failed to retrieve run for algorithm %s with run id %s: %v", algorithm, id, err)
	}

	fmt.Println("Response:", response)
	return nil
}
