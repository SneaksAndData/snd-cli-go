package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

func NewCmdGetPayload(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "payload",
		Short:   heredoc.Doc(`Get the input payload for a Nexus run`),
		Example: heredoc.Doc(`snd nx payload --id 762b07c-c67a-4327-970a-18d923fd --algorithm rdc-auto-replenishment-crystal-orchestrator -e production`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("nx", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := executeGetPayload(service.(*cmdutil.NexusService), id, template)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Nexus run identifier")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("failed to mark 'id' as a required flag: %w", err)
		return nil
	}
	return cmd
}

func executeGetPayload(nexus *cmdutil.NexusService, id, template string) (string, error) {
	err := nexus.Authenticate()
	if err != nil {
		return "", err
	}

	response, err := nexus.Client.GetRunPayload(id, template)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve payload for algorithm %s with run id %s: %w", template, id, err)
	}

	return response, nil
}
