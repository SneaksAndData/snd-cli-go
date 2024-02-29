package ml

import (
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

var payload, tag string

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("algorithm", env, authService)
			if err != nil {
				log.Fatalf(err.Error())
			}
			payloadPath := file.File{FilePath: payload}
			resp, err := runRun(service.(*algorithmClient.Service), payloadPath, algorithm, tag)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "Client-side submission identifier")
	cmd.MarkFlagRequired("payload")
	return cmd
}

func runRun(algorithmService Service, fileOp Operations, algorithm, tag string) (string, error) {
	payloadJSON, err := fileOp.ReadJSONFile()
	if err != nil {
		return "", err
	}
	response, err := algorithmService.CreateRun(algorithm, payloadJSON, tag)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}

	return response, nil
}
