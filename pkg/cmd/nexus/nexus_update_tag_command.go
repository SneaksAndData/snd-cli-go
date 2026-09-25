package nexus

import (
	"fmt"
	"snd-cli/pkg/cmdutil"

	"github.com/MakeNowJust/heredoc"
	api "github.com/SneaksAndData/nexus-sdk-go/pkg/generated/scheduler"
	"github.com/SneaksAndData/nexus-sdk-go/sdk"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewCmdUpdateTag(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	var newTag string
	cmd := &cobra.Command{
		Use:   "run",
		Short: heredoc.Doc(`Update a client tag on a completed Nexus run.`),
		Long:  heredoc.Doc(`Update a client tag to a provided value. Clients that rely on tagging can utilize this to reset their internal run tracking state`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("nx", env, url, authService)
			if err != nil {
				return err
			}
			err = executeTagUpdate(service.(*cmdutil.NexusService), id, template, newTag)
			if err == nil {
				pterm.DefaultBasicText.Println(fmt.Sprintf("Submission %s/%s tag successfully updated to %s", template, id, newTag))
			}
			return err
		},
		Example: heredoc.Doc(`
 $ snd nx tag --template my-algorithm --id 762b07c-c67a-4327-970a-18d923fd --value my-new-tag
}
`),
	}

	cmd.Flags().StringVarP(&newTag, "value", "v", "", "New tag to assign to a submission")

	err := cmd.MarkFlagRequired("value")
	if err != nil {
		fmt.Println("failed to mark 'value' as a required flag: %w", err)
		return nil
	}

	return cmd
}

func executeTagUpdate(nexus *cmdutil.NexusService, tag, template, id string) error {
	err := nexus.Authenticate()
	if err != nil {
		return err
	}

	response, err := nexus.Client.GetRun(id, template)
	if err != nil {
		return fmt.Errorf("failed to retrieve run for the template %s with run id %s: %w", template, id, err)
	}

	if !sdk.IsFinished(response.Status.Value) {
		return fmt.Errorf("cannot update tag for an in-progress run %s/%s (%s)", template, id, response.Status.Value)
	}

	err = nexus.Client.UpdateRunTag(&api.ModelsTagUpdateRequest{NewTag: api.OptString{Set: true, Value: tag}}, id, template)

	if err != nil {
		return fmt.Errorf("failed to update tag for %s/%s: %w", template, id, err)
	}

	return nil
}
