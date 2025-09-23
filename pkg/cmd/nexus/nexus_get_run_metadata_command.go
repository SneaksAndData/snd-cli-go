package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

func NewCmdGetRunMetadata(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "meta",
		Short:   heredoc.Doc(`Get metadata for a Nexus run`),
		Example: heredoc.Doc(`snd nx meta --id 762b07c-c67a-4327-970a-18d923fd --algorithm rdc-auto-replenishment-crystal-orchestrator -e production`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("nx", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := executeGetMetadata(service.(*cmdutil.NexusService), id, template)
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

func executeGetMetadata(nexus *cmdutil.NexusService, id, template string) (string, error) {
	err := nexus.Authenticate()
	if err != nil {
		return "", err
	}
	response, err := nexus.Client.GetMetadata(id, template)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve metadata for the template %s with run id %s: %w", template, id, err)
	}

	serialized, err := response.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("failed to serialize metadata response for the template %s with run id %s: %w", template, id, err)
	}

	return string(serialized), nil
}
