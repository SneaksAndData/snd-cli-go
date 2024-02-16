package ml

import (
	"fmt"
	algorithms "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util"
)

var payload, tag string

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {

			var algorithmService, err = InitAlgorithmService(fmt.Sprintf(crystalBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return runRun(algorithmService)
		},
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "Client-side submission identifier")
	return cmd
}

func runRun(algorithmService *algorithms.Service) error {
	payloadJSON, err := util.ReadJSONFile(payload)
	if err != nil {
		log.Fatal(err)
	}
	response, err := algorithmService.CreateRun(algorithm, payloadJSON, tag)
	if err != nil {
		log.Fatalf("Failed to retrieve run for algorithm %s with run id %s: %v", algorithm, id, err)
	}

	fmt.Println("Response:", response)
	return nil
}
