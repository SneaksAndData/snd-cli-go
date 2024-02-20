package ml

import (
	"fmt"
	algorithms "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmd/util/token"
)

const crystalBaseURL = "https://crystal.%s.sneaksanddata.com"

var env, authProvider, algorithm string

type Service interface {
	RetrieveRun(runID string, algorithmName string) (string, error)
	CreateRun(algorithmName string, input map[string]interface{}, tag string) (string, error)
}

type Operations interface {
	ReadJSONFile() (map[string]interface{}, error)
}

type FileServiceFactory func(path string) (file.File, error)

func NewCmdAlgorithm() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "algorithm",
		Short:   "Manage ML algorithm jobs",
		GroupID: "ml",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "l", "", "Specify the algorithm name")

	var algorithmService, err = InitAlgorithmService(fmt.Sprintf(crystalBaseURL, env))
	if err != nil {
		log.Fatalf("Failed to initialize algorithm service: %v", err)
	}
	cmd.AddCommand(NewCmdGet(algorithmService))
	cmd.AddCommand(NewCmdRun(algorithmService, func(path string) (file.File, error) {
		// This anonymous function acts as the factory. It encapsulates
		// the logic to create a new claim service instance.
		return file.File{FilePath: path}, nil
	}))
	return cmd
}

func InitAlgorithmService(url string) (*algorithms.Service, error) {
	tc := token.TokenCache{}
	config := algorithms.Config{
		GetTokenFunc: tc.ReadToken,
		SchedulerURL: url,
		APIVersion:   "v1.2",
	}

	algorithmService, err := algorithms.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create algorithm service: %v", err)
	}
	return algorithmService, nil
}
