package ml

import (
	"encoding/json"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

var payload, tag string

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		Short: heredoc.Doc(`Run a ML Algorithm.

The payload should be provided as a JSON file with the structure below.

<pre><code>
{
 "AlgorithmName": "<string> (optional) - The name of the algorithm to run",
 "AlgorithmParameters": "<object> (required) - Any additional parameters for the algorithm",
 "CustomConfiguration": {
	"imageRepository": "",
    "imageTag": "",
    "deadlineSeconds": 123,
    "maximumRetries": 1,
    "cpuLimit": "",
    "memoryLimit": "",
    "monitoringParameters": [],
    "speculativeAttempts": 2,
},  - <CustomConfiguration> (optional) - Custom configuration for the algorithm",
 "Tag": "<string> (optional) - Client-side submission identifier"
}
</code></pre>
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			payloadPath := file.File{FilePath: payload}
			resp, err := runRun(service.(*algorithmClient.Service), payloadPath, algorithm, tag)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
		Example: heredoc.Doc(`snd algorithm run --algorithm rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json`),
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "Client-side submission identifier")
	return cmd
}

func runRun(algorithmService Service, fileOp Operations, algorithm, tag string) (string, error) {
	p, err := readAlgorithmPayload(fileOp)
	if err != nil {
		return "", err
	}
	response, err := algorithmService.CreateRun(algorithm, p, tag)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}

	return response, nil
}

func readAlgorithmPayload(fileOp Operations) (algorithmClient.Payload, error) {
	var p = algorithmClient.Payload{
		AlgorithmParameters: nil,
		AlgorithmName:       "",
		CustomConfiguration: algorithmClient.CustomConfiguration{},
		Tag:                 "",
	}
	if fileOp.IsValidPath() {
		content, err := fileOp.ReadJSONFile()
		if err != nil {
			return p, fmt.Errorf("failed to read JSON file: %w", err)
		}
		var payload *algorithmClient.Payload
		c, err := json.Marshal(content)
		if err != nil {
			return p, fmt.Errorf("error marshaling content from file: %w", err)
		}
		if err = json.Unmarshal(c, &payload); err != nil {
			return p, fmt.Errorf("error unmarshaling content to algorithm.Payload: %w", err)
		}
		return *payload, nil
	}
	return p, fmt.Errorf("payload path is not valid")
}
