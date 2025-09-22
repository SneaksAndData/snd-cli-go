package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	api "github.com/SneaksAndData/nexus-sdk-go/pkg/generated/scheduler"
	"github.com/go-faster/jx"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmdutil"
)

// CommandConfig holds the configuration for the run command.
type CommandConfig struct {
	Payload string
}

func NewCmdRun(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	var config CommandConfig
	cmd := &cobra.Command{
		Use:   "run",
		Short: heredoc.Doc(`Run a Nexus algorithm. The payload should be provided as a JSON file with the structure below. If no payload is needed, add {} to file contents`),
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := executeCreate(config, authServiceFactory, serviceFactory)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
		Example: heredoc.Doc(`snd nx run --template rdc-auto-replenishment-crystal-orchestrator --payload /path/to/payload.json`),
	}

	cmd.Flags().StringVarP(&config.Payload, "payload", "p", "", "Path to the payload JSON file")

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

	resp, err := createRun(service.(*cmdutil.NexusService), config.Payload, template)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// createRun runs the algorithm service with the provided parameters.
func createRun(nexus *cmdutil.NexusService, payloadPath, template string) (string, error) {
	payload, err := readRunPayload(payloadPath)

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

	return response, nil
}

// readRunPayload reads and unmarshal the algorithm payload from the provided path.
func readRunPayload(payloadPath string) (*api.ModelsAlgorithmRequest, error) {
	defaultTag, _ := util.GenerateTag()
	if payloadPath == "" {
		pterm.DefaultBasicText.Println(pterm.Sprintf("You have not provided a payload path, thus a client tag for this submission is not available. Using generated tag: %s \n", defaultTag))
		return &api.ModelsAlgorithmRequest{
			AlgorithmParameters: map[string]jx.Raw{},
			CustomConfiguration: api.OptV1NexusAlgorithmSpec{
				Set: false,
			},
			ParentRequest: api.OptModelsAlgorithmRequestRef{
				Set: false,
			},
			PayloadValidFor: api.OptString{
				Set: false,
			},
			RequestApiVersion: api.OptString{
				Set: false,
			},
			Tag: api.OptString{
				Set:   true,
				Value: defaultTag,
			},
		}, nil
	}

	payloadData, err := os.ReadFile(payloadPath)

	if err != nil {
		return nil, err
	}

	userPayload := &api.ModelsAlgorithmRequest{}
	err = userPayload.UnmarshalJSON(payloadData)
	if err != nil {
		return nil, err
	}

	if !userPayload.Tag.Set {
		pterm.DefaultBasicText.Println(pterm.Sprintf("Payload does not provide a client tag. Using generated tag: %s \n", defaultTag))
		userPayload.Tag.Set = true
		userPayload.Tag.Value = defaultTag
	}

	return userPayload, nil
}
