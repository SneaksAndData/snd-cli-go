package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var id string

func NewCmdGet(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   heredoc.Doc(`Get the result for a Nexus run`),
		Example: heredoc.Doc(`snd nx get --id 762b07c-c67a-4327-970a-18d923fd --template omni-channel-solver -e production`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("nx", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := getRun(service.(*cmdutil.NexusService), id, template)
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

func getRun(nexus *cmdutil.NexusService, id, template string) (string, error) {
	response, err := nexus.Client.GetRun(id, template)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve run for the template %s with run id %s: %w", template, id, err)
	}

	serialized, err := response.MarshalJSON()

	if err != nil {
		return "", fmt.Errorf("failed to serialize response for the template %s with run id %s: %w", template, id, err)
	}

	return string(serialized), nil
}
