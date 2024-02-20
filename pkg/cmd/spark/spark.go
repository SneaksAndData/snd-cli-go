package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util/token"
)

const beastBaseURL = "https://beast-v3.%s.sneaksanddata.com"

var env, authProvider, id string

type Service interface {
	GetConfiguration(name string) (spark.SubmissionConfiguration, error)
	GetLogs(id string) (string, error)
	GetLifecycleStage(id string) (interface{}, error)
	GetRuntimeInfo(id string) (string, error)
	RunJob(request spark.JobParams, sparkJobName string) (string, error)
}

func NewCmdSpark() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spark",
		Short:   "Manage Spark jobs",
		GroupID: "spark",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "Specify the  Job ID")

	var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
	if err != nil {
		log.Fatalf("Failed to initialize spark service: %v", err)
	}

	cmd.AddCommand(NewCmdSubmit(sparkService))
	cmd.AddCommand(NewCmdRuntimeInfo(sparkService))
	cmd.AddCommand(NewCmdRequestStatus(sparkService))
	cmd.AddCommand(NewCmdLogs(sparkService))
	cmd.AddCommand(NewCmdConfiguration(sparkService))
	cmd.AddCommand(NewCmdEncrypt(sparkService))

	return cmd
}

func InitSparkService(url string) (*spark.Service, error) {
	tc := token.TokenCache{}
	config := spark.Config{
		BaseURL:      url,
		GetTokenFunc: tc.ReadToken,
	}

	sparkService, err := spark.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create spark service: %v", err)
	}
	return sparkService, nil

}
