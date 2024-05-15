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
		return "", fmt.Errorf("failed to create run for algorithm %s: %w", algorithm, err)
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
		c, err := json.Marshal(content)
		if err != nil {
			return p, fmt.Errorf("error marshaling content from file: %w", err)
		}
		var userPayload *Payload
		if err = json.Unmarshal(c, &userPayload); err != nil {
			return p, fmt.Errorf("error unmarshaling content to algorithm.userPayload: %w", err)
		}

		crystalPayload, _ := json.Marshal(userPayload)
		fmt.Println(string(crystalPayload))

		var payload *algorithmClient.Payload

		if err = json.Unmarshal(crystalPayload, &payload); err != nil {
			return p, fmt.Errorf("error unmarshaling content to algorithm.Payload: %w", err)
		}
		return *payload, nil
	}
	return p, fmt.Errorf("payload path is not valid")
}

// Payload defines the structure of the request body for creating algorithm runs.
type Payload struct {
	AlgorithmParameters map[string]interface{} `validate:"required" json:"algorithm_parameters"`
	AlgorithmName       string                 `json:"algorithm_name"`
	CustomConfiguration CustomConfiguration    `json:"custom_configuration"`
	Tag                 string                 `json:"tag"`
}

func (p Payload) MarshalJSON() ([]byte, error) {
	type T struct {
		AlgorithmParameters map[string]interface{} `validate:"required" json:"AlgorithmParameters"`
		AlgorithmName       string                 `json:"AlgorithmName"`
		CustomConfiguration CustomConfiguration    `json:"CustomConfiguration"`
		Tag                 string                 `json:"Tag"`
	}

	return json.Marshal(T(p))
}

type CustomConfiguration struct {
	ImageRepository      *string              `json:"image_repository"`
	ImageTag             *string              `json:"image_tag"`
	DeadlineSeconds      *int                 `json:"deadline_seconds"`
	MaximumRetries       *int                 `json:"maximum_retries"`
	Env                  []ConfigurationEntry `json:"env"`
	Secrets              []string             `json:"secrets"`
	Args                 []ConfigurationEntry `json:"args"`
	CpuLimit             *string              `json:"cpu_limit"`
	MemoryLimit          *string              `json:"memory_limit"`
	Workgroup            *string              `json:"workgroup"`
	AdditionalWorkgroups map[string]string    `json:"additional_workgroups"`
	Version              *string              `json:"version"`
	MonitoringParameters []string             `json:"monitoring_parameters"`
	CustomResources      map[string]string    `json:"custom_resources"`
	SpeculativeAttempts  *int                 `json:"speculative_attempts"`
}

func (c CustomConfiguration) MarshalJSON() ([]byte, error) {
	type T struct {
		ImageRepository      *string              `json:"imageRepository"`
		ImageTag             *string              `json:"imageTag"`
		DeadlineSeconds      *int                 `json:"deadlineSeconds"`
		MaximumRetries       *int                 `json:"maximumRetries"`
		Env                  []ConfigurationEntry `json:"env"`
		Secrets              []string             `json:"secrets"`
		Args                 []ConfigurationEntry `json:"args"`
		CpuLimit             *string              `json:"cpuLimit"`
		MemoryLimit          *string              `json:"memoryLimit"`
		Workgroup            *string              `json:"workgroup"`
		AdditionalWorkgroups map[string]string    `json:"additionalWorkgroups"`
		Version              *string              `json:"version"`
		MonitoringParameters []string             `json:"monitoringParameters"`
		CustomResources      map[string]string    `json:"customResources"`
		SpeculativeAttempts  *int                 `json:"speculativeAttempts"`
	}

	return json.Marshal(T(c))
}

type ConfigurationEntry struct {
	Name      string                                  `json:"name"`
	Value     string                                  `json:"value"`
	ValueType *algorithmClient.ConfigurationValueType `json:"value_from"`
}

func (c ConfigurationEntry) MarshalJSON() ([]byte, error) {
	type T struct {
		Name      string                                  `json:"name"`
		Value     string                                  `json:"value"`
		ValueType *algorithmClient.ConfigurationValueType `json:"valueFrom"`
	}

	return json.Marshal(T(c))
}
