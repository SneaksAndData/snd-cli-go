package nexus

import (
	"errors"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	api "github.com/SneaksAndData/nexus-sdk-go/pkg/generated/scheduler"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmdutil"
)

// CommandConfig holds the configuration for the run command.
type CommandConfig struct {
	PayloadPath          string
	CustomConfigPath     string
	ParentRequestRefPath string
	ValidFor             string
	Tag                  string
}

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	var config CommandConfig
	cmd := &cobra.Command{
		Use:   "run",
		Short: heredoc.Doc(`Run a Nexus algorithm. The payload should be provided as a valid JSON file. In addition, configuration overrides, parent reference and validity period could be overridden.`),
		Long: heredoc.Doc(`Run a Nexus algorithm. The payload should be provided as a valid JSON file. In addition, configuration overrides, parent reference and validity period could be overridden.	
Custom configuration format is provided here: https://github.com/SneaksAndData/nexus-sdk-go/blob/main/pkg/generated/scheduler/oas_schemas_gen.go#L1062-L1069. Example for a version override: {"container": {"versionTag": "v1.2.0"}}
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := executeCreate(config, authServiceFactory, serviceFactory)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
		Example: heredoc.Doc(`
 $ snd nx run --template rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json
 $ snd nx run --template rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json --custom-configuration /path/to/config.json
 $ snd nx run --template rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json --custom-configuration /path/to/config.json --parent-request /path/to/request.json --valid-for 168h --tag mytag
}
`),
	}

	cmd.Flags().StringVarP(&config.PayloadPath, "payload", "p", "", "Path to the input payload (json).")
	cmd.Flags().StringVar(&config.CustomConfigPath, "custom-configuration", "", "Path to the optional custom configuration for the run (json).")
	cmd.Flags().StringVar(&config.ParentRequestRefPath, "parent-request", "", "Path to the optional parent request reference (json).")
	cmd.Flags().StringVar(&config.ValidFor, "valid-for", "24h", "Payload validity period override in hours. Defaults to 24h, maximum possible value is 168h.")
	cmd.Flags().StringVar(&config.Tag, "tag", "", "Client-side tag for identifying the submission.")

	err := cmd.MarkFlagRequired("payload")
	if err != nil {
		fmt.Println("failed to mark 'payload' as a required flag: %w", err)
		return nil
	}

	return cmd
}

func executeCreate(config CommandConfig, authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) (string, error) {
	authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
	if err != nil {
		return "", err
	}

	if util.IsProdEnv(env) {
		answer := util.InteractiveContinue()
		if answer != "yes" {
			return "", fmt.Errorf("operation aborted by user")
		}
	}

	service, err := serviceFactory.CreateService("nx", env, url, authService)
	if err != nil {
		return "", err
	}

	resp, err := createRun(service.(*cmdutil.NexusService), config, template)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// createRun runs the algorithm service with the provided parameters.
func createRun(nexus *cmdutil.NexusService, runConfig CommandConfig, template string) (string, error) {
	payload, err := generateRequest(runConfig)

	if err != nil {
		return "", err
	}

	err = nexus.Authenticate()
	if err != nil {
		return "", err
	}

	response, err := nexus.Client.CreateRun(payload, template, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create run for the template %s: %w", template, err)
	}

	return fmt.Sprintf("{\"requestId\": \"%s\"}", response), nil
}

// generateRequest reads and unmarshal the algorithm payload from the provided path.
func generateRequest(runConfig CommandConfig) (*api.ModelsAlgorithmRequest, error) {
	var payloadData, customConfigData, parentRequestData []byte
	var tag string
	algParams := &api.ModelsAlgorithmRequestAlgorithmParameters{}
	customConfig := &api.V1NexusAlgorithmSpec{}
	parentRef := &api.ModelsAlgorithmRequestRef{}

	if runConfig.PayloadPath == "" {
		return nil, errors.New("payload path is an empty string - cannot generate a run configuration")
	}

	if runConfig.Tag == "" {
		tag, _ = util.GenerateTag()
		pterm.DefaultBasicText.Println(pterm.Sprintf("Client tag not provided, using generated tag: %s \n", tag))
	} else {
		tag = runConfig.Tag
	}
	// read and parse data payload first
	payloadData, err := os.ReadFile(runConfig.PayloadPath)

	if err != nil {
		return nil, err
	}

	err = algParams.UnmarshalJSON(payloadData)
	if err != nil {
		return nil, err
	}

	// load custom config if provided
	if runConfig.CustomConfigPath != "" {
		customConfigData, err = os.ReadFile(runConfig.CustomConfigPath)
		if err != nil {
			return nil, err
		}
		err = customConfig.UnmarshalJSON(customConfigData)
		if err != nil {
			return nil, err
		}
	}

	if runConfig.ParentRequestRefPath != "" {
		parentRequestData, err = os.ReadFile(runConfig.ParentRequestRefPath)
		if err != nil {
			return nil, err
		}
		err = parentRef.UnmarshalJSON(parentRequestData)
		if err != nil {
			return nil, err
		}
	}

	runPayload := &api.ModelsAlgorithmRequest{
		AlgorithmParameters: *algParams,
		CustomConfiguration: api.OptV1NexusAlgorithmSpec{
			Set:   runConfig.CustomConfigPath != "",
			Value: *customConfig,
		},
		ParentRequest: api.OptModelsAlgorithmRequestRef{
			Set:   runConfig.ParentRequestRefPath != "",
			Value: *parentRef,
		},
		PayloadValidFor: api.OptString{
			Set:   true,
			Value: runConfig.ValidFor,
		},
		RequestApiVersion: api.OptString{
			Set: false,
		},
		Tag: api.OptString{
			Set:   true,
			Value: tag,
		},
	}

	return runPayload, nil
}
