package ml

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var id string

func NewCmdGet(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   heredoc.Doc(`Get the result for a ML Algorithm run`),
		Example: heredoc.Doc(`snd algorithm get --id 762b07c-c67a-4327-970a-18d923fd --algorithm rdc-auto-replenishment-crystal-orchestrator -e production`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := getRun(service.(*algorithmClient.Service), id, algorithm)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("failed to mark 'id' as a required flag: %w", err)
		return nil
	}
	return cmd
}

func getRun(algorithmService Service, id, algorithm string) (string, error) {
	response, err := algorithmService.RetrieveRun(id, algorithm)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to find run for algorithm %s with run id %s : %v", algorithm, id, "Run not found")
		}
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}

	prettifyResponse, err := util.PrettifyJSON(response)
	if err != nil {
		return "", fmt.Errorf("failed to prettify response: %w", err)
	}

	return prettifyResponse, nil
}
