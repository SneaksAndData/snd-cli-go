package ml

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

const crystalURL = "https://crystal.%s.sneaksanddata.com"

var env, url, authProvider, algorithm string

type Service interface {
	RetrieveRun(runID string, algorithmName string) (string, error)
	CreateRun(algorithmName string, input algorithmClient.Payload, tag string) (string, error)
	CancelRun(algorithmName string, requestId string, initiator string, reason string) (string, error)
}

//type Operations interface {
//	ReadJSONFile() (map[string]interface{}, error)
//}

type FileServiceFactory func(path string) (file.File, error)

func NewCmdAlgorithm(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "algorithm",
		Short: "Manage ML algorithm jobs",
		Long:  "Manage ML algorithm jobs",
		Example: heredoc.Doc(`
			$ snd algorithm run --algorithm store-auto-replenishment-crystal-orchestrator --payload ./crystal-payload.json
			$ snd algorithm get --id fa1d02af-c294-4bf6-989f-1234 --algorithm store-auto-replenishment-crystal-orchestrator
		`),
		GroupID: "ml",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "", "", "Specify the algorithm name")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", crystalURL, "Specify the service url")

	err := cmd.MarkPersistentFlagRequired("algorithm")
	if err != nil {
		fmt.Println("failed to mark 'algorithm' as a required flag: %w", err)
		return nil
	}

	cmd.AddCommand(NewCmdGet(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRun(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdCancel(authServiceFactory, serviceFactory))
	return cmd
}
