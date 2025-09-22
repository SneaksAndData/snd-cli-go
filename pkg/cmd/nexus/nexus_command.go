package nexus

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

const nexusServiceUrl = "https://nexus.%s.sneaksanddata.com"

var env, url, authProvider, template, authUrl string

func NewCmdNexus(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nx",
		Short: "Manage Nexus runs",
		Long:  "Operate Nexus runs from your command line: create, read results and more!",
		Example: heredoc.Doc(`
			$ snd nx run --template omni-channel-solver --payload ./payload.json
			$ snd nx get --id fa1d02af-c294-4bf6-989f-1234 --template omni-channel-solver
			$ snd nx payload --id fa1d02af-c294-4bf6-989f-1234 --template omni-channel-solver
			$ snd nx cancel --id fa1d02af-c294-4bf6-989f-1234 --template omni-channel-solver --reason test
		`),
		GroupID: "nx",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", cmdutil.BaseEnvironment, "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&template, "template", "", "", "Specify the template name")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", nexusServiceUrl, "Specify the service url")
	cmd.PersistentFlags().StringVarP(&authUrl, "custom-auth-url", "", "", "Specify the auth service uri")

	err := cmd.MarkPersistentFlagRequired("template")
	if err != nil {
		fmt.Println("failed to mark 'nx' as a required flag: %w", err)
		return nil
	}

	cmd.AddCommand(NewCmdGet(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRun(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdCancel(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdGetPayload(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdGetRunMetadata(authServiceFactory, serviceFactory))
	return cmd
}
