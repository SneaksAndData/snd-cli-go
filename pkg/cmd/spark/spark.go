package spark

import (
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util"
)

var env, authProvider, id string

func NewCmdSpark() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spark",
		Short:   "Manage Spark jobs",
		GroupID: "spark",
	}

	cmd.AddCommand(NewCmdSubmit())
	cmd.AddCommand(NewCmdRuntimeInfo())
	cmd.AddCommand(NewCmdRequestStatus())
	cmd.AddCommand(NewCmdLogs())
	cmd.AddCommand(NewCmdConfiguration())
	cmd.AddCommand(NewCmdEncrypt())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "Specify the  Job ID")

	return cmd
}

func InitSparkService(url string) (*spark.Service, error) {
	config := spark.Config{
		BaseURL:      url,
		GetTokenFunc: util.ReadToken,
	}

	sparkService, err := spark.New(config)
	if err != nil {
		log.Fatalf("Failed to create spark service: %v", err)
	}
	return sparkService, nil

}
