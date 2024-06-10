package ml

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

// CommandConfig holds the configuration for the run command.
type CommandConfig struct {
	Payload string
	Tag     string
}

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	var config CommandConfig
	cmd := &cobra.Command{
		Use: "run",
		Short: heredoc.Doc(`Run a ML Algorithm.

The payload should be provided as a JSON file with the structure below.

<pre><code>
{
 "algorithm_name": "&lt;string&gt;  (optional) - The name of the algorithm to run",
 "algorithm_parameters": "&lt;object&gt; (required) - Any additional parameters for the algorithm",
 "custom_configuration": {
	"image_repository": &lt;string&gt;,
	"image_tag": &lt;string&gt;,
	"deadline_seconds": &lt;int&gt;,
	"maximum_retries": &lt;int&gt;,
	"env": {"name": &lt;string&gt;, "value": &lt;string&gt;, "value_from": "PLAIN" | "RELATIVE_REFERENCE"}
	"secrets":  &lt;string[]&gt;,
	"args": {"name": &lt;string&gt;, "value": &lt;string&gt;, "value_from": "PLAIN" | "RELATIVE_REFERENCE"},
	"cpu_limit": &lt;string&gt;,
	"memory_limit": &lt;string&gt;,
	"workgroup": &lt;string&gt;,
	"additional_workgroups": &lt;map[string]string&gt;,
	"version": &lt;string&gt;,
	"monitoring_parameters": &lt;string[]&gt;,
	"custom_resources": &lt;map[string]string&gt;,
	"speculative_attempts": int
} - &lt;CustomConfiguration&gt; (optional) - Custom configuration for the algorithm",
 "tag": "&lt;string&gt; (optional) - Client-side submission identifier"
}
</code></pre>
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRun(config, authServiceFactory, serviceFactory)
		},
		Example: heredoc.Doc(`snd algorithm run --algorithm rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json`),
	}

	cmd.Flags().StringVarP(&config.Payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&config.Tag, "tag", "t", "", "Client-side submission identifier")

	err := cmd.MarkFlagRequired("payload")
	if err != nil {
		fmt.Println("failed to mark 'payload' as a required flag: %w", err)
		return nil
	}

	return cmd
}

func runRun(config CommandConfig, authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) error {
	authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
	if err != nil {
		return err
	}

	service, err := serviceFactory.CreateService("algorithm", env, url, authService)
	if err != nil {
		return err
	}

	resp, err := runAlgorithm(service.(*algorithmClient.Service), config.Payload, algorithm, config.Tag)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

// runAlgorithm runs the algorithm service with the provided parameters.
func runAlgorithm(algorithmService Service, payloadPath, algorithm, tag string) (string, error) {
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

// readAlgorithmPayload reads and unmarshal the algorithm payload from the provided path.
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

	var userPayload *Payload
	err := f.ReadAndUnmarshal(&userPayload)
	if err != nil {
		return p, err
	}

	var payload algorithmClient.Payload
	err = util.ConvertStruct(*userPayload, &payload)
	if err != nil {
		return p, err
	}

	return payload, nil

}
