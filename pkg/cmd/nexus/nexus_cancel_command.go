package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	api "github.com/SneaksAndData/nexus-sdk-go/pkg/generated/scheduler"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/token"
	"snd-cli/pkg/cmdutil"
)

var reason string

func NewCmdCancel(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel",
		Short:   heredoc.Doc(`Cancel a Nexus run`),
		Example: heredoc.Doc(`snd nx cancel --id 762b07c-c67a-4327-970a-18d923fd --template rdc-auto-replenishment-crystal-orchestrator -e production`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			tokenProvider, err := token.NewProvider(authService, env)
			if err != nil {
				return fmt.Errorf("unable to create token provider: %w", err)
			}

			user := tokenProvider.GetUserFromToken()

			service, err := serviceFactory.CreateService("nx", env, url, authService)
			if err != nil {
				return err
			}

			resp, err := cancelRun(service.(*cmdutil.NexusService), template, id, user, reason)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Nexus run identifier")
	cmd.Flags().StringVarP(&reason, "reason", "", "", "Specify reason for cancelling the run")

	err := cmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("failed to mark 'id' as a required flag: %w", err)
		return nil
	}

	err = cmd.MarkFlagRequired("reason")
	if err != nil {
		fmt.Println("failed to mark 'reason' as a required flag: %w", err)
		return nil
	}

	return cmd
}

func cancelRun(nexus *cmdutil.NexusService, template, id, initiator, reason string) (string, error) {
	err := nexus.Authenticate()
	if err != nil {
		return "", err
	}

	err = nexus.Client.CancelRun(&api.ModelsCancellationRequest{
		CancellationPolicy: api.OptString{
			Set:   true,
			Value: "Background",
		},
		Initiator: api.OptString{
			Set:   true,
			Value: initiator,
		},
		Reason: api.OptString{
			Set:   true,
			Value: reason,
		},
	}, id, template)

	if err != nil {
		return "", fmt.Errorf("failed to cancel run for the template %s with run id %s: %w", template, id, err)
	}

	return fmt.Sprintf("Cancel initiated by %s for: %s/%s, reason: %s", initiator, template, id, reason), nil
}
