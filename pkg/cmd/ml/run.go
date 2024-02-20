package ml

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var payload, tag string

func NewCmdRun(service Service, factory FileServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileService, err := factory(payload)
			if err != nil {
				log.Fatalf(err.Error())
			}
			resp, err := runRun(service, fileService, algorithm, tag)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "Client-side submission identifier")
	return cmd
}

func runRun(algorithmService Service, fileOp Operations, algorithm, tag string) (string, error) {
	payloadJSON, err := fileOp.ReadJSONFile()
	if err != nil {
		return "", err
	}
	response, err := algorithmService.CreateRun(algorithm, payloadJSON, tag)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %v", algorithm, id, err)
	}

	return response, nil
}
