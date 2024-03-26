package ml

import (
	"encoding/json"
	"errors"
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

var payload, tag string

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := runRun(service.(*algorithmClient.Service), payload, algorithm, tag)
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

func runRun(algorithmService Service, payloadPath string, algorithm, tag string) (string, error) {
	p, err := readAlgorithmPayload(payloadPath)
	if err != nil {
		return "", err
	}
	response, err := algorithmService.CreateRun(algorithm, p, tag)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}

	return response, nil
}

func readAlgorithmPayload(path string) (algorithmClient.Payload, error) {
	var p = algorithmClient.Payload{
		AlgorithmParameters: nil,
		AlgorithmName:       "",
		CustomConfiguration: algorithmClient.CustomConfiguration{},
		Tag:                 "",
	}
	if path == "" {
		return p, errors.New("no payload path provided")
	}
	f := file.File{FilePath: path}
	if f.IsValidPath() {
		content, err := f.ReadJSONFile()
		if err != nil {
			return p, fmt.Errorf("failed to read JSON file '%s': %w", path, err)
		}
		var payload *algorithmClient.Payload
		c, err := json.Marshal(content)
		if err != nil {
			return p, fmt.Errorf("error marshaling content from file '%s': %w", path, err)
		}
		if err = json.Unmarshal(c, &payload); err != nil {
			return p, fmt.Errorf("error unmarshaling content to algorithm.Payload: %w", err)
		}
		return *payload, nil
	}
	return p, errors.New("payload path is not valid")
}
