package ml

import (
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
)

const crystalURL = "https://crystal.%s.sneaksanddata.com"

var env, url, authProvider, algorithm string

type Service interface {
	RetrieveRun(runID string, algorithmName string) (string, error)
	CreateRun(algorithmName string, input map[string]interface{}, tag string) (string, error)
}

type Operations interface {
	ReadJSONFile() (map[string]interface{}, error)
}

type FileServiceFactory func(path string) (file.File, error)

func NewCmdAlgorithm(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "algorithm",
		Short:   "Manage ML algorithm jobs",
		GroupID: "ml",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "", "", "Specify the algorithm name")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", crystalURL, "Specify the service url")

	cmd.AddCommand(NewCmdGet(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRun(authServiceFactory, serviceFactory))
	return cmd
}
