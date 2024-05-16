package ml

import (
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
 "algorithm_name": "<i>&lt;string&gt;</i>  (optional) - The name of the algorithm to run",
 "algorithm_parameters": "<object> (required) - Any additional parameters for the algorithm",
 "custom_configuration": {
	"image_repository": <i>&lt;string&gt;</i>,
    "image_tag": <i>&lt;string&gt;</i>,
    "deadline_seconds": <i>&lt;int&gt;</i>,
    "maximum_retries": <i>&lt;int&gt;</i>,
	"env": {"name": <i>&lt;string&gt;</i>, "value": <i>&lt;string&gt;</i>, "value_from": "PLAIN" | "RELATIVE_REFERENCE"}
	"secrets": <string[]>,
	"args": {"name": <i>&lt;string&gt;</i>, "value": <i>&lt;string&gt;</i>, "value_from": "PLAIN" | "RELATIVE_REFERENCE"},
    "cpu_limit": <i>&lt;string&gt;</i>,
    "memory_limit": <i>&lt;string&gt;</i>,
	"workgroup": <i>&lt;string&gt;</i>,
	"additional_workgroups": <map[string]string>,
	"version": <i>&lt;string&gt;</i>,
    "monitoring_parameters": <i>&lt;[]string&gt;</i>,
	"custom_resources": <i>&lt;map[string]string&gt;</i>,
    "speculative_attempts": int,},- <CustomConfiguration> (optional) - Custom configuration for the algorithm",
 "tag": "<i>&lt;string&gt;</i> (optional) - Client-side submission identifier"
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
			resp, err := runRun(service.(*algorithmClient.Service), payload, algorithm, tag)
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

func runRun(algorithmService Service, payloadPath, algorithm, tag string) (string, error) {
	p, err := readAlgorithmPayload(payloadPath)
	if err != nil {
		return "", err
	}
	response, err := algorithmService.CreateRun(algorithm, p, tag)
	if err != nil {
		return "", fmt.Errorf("failed to create run for algorithm %s: %w", algorithm, err)
	}

	return response, nil
}

func readAlgorithmPayload(payloadPath string) (algorithmClient.Payload, error) {
	var p = algorithmClient.Payload{
		AlgorithmParameters: nil,
		AlgorithmName:       "",
		CustomConfiguration: algorithmClient.CustomConfiguration{},
		Tag:                 "",
	}
	if payloadPath == "" {
		return p, nil
	}
	f := file.File{FilePath: payloadPath}
	var payload algorithmClient.Payload
	err := f.ReadAndUnmarshal(&payload)
	if err != nil {
		return p, fmt.Errorf("error unmarshaling content to algorithm.Payload: %w", err)
	}
	return payload, nil
}
