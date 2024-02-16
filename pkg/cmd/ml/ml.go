package ml

import (
	algorithms "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util"
)

var env, authProvider, algorithm string

func NewCmdAlgorithm() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "algorithm",
		Short:   "Manage ML algorithm jobs",
		GroupID: "ml",
	}

	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdRun())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "l", "", "Specify the algorithm name")

	return cmd
}

func InitAlgorithmService(url string) (*algorithms.Service, error) {
	config := algorithms.Config{
		GetTokenFunc: util.ReadToken,
		SchedulerURL: url,
		APIVersion:   "v1.2",
	}

	algorithmService, err := algorithms.New(config)
	if err != nil {
		log.Fatalf("Failed to create algorithm service: %v", err)
	}
	return algorithmService, nil
}
