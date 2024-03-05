package spark

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

const beastURL = "https://beast-v3.%s.sneaksanddata.com"

var env, url, authProvider, id string

type Service interface {
	GetConfiguration(name string) (spark.SubmissionConfiguration, error)
	GetLogs(id string) (string, error)
	GetLifecycleStage(id string) (interface{}, error)
	GetRuntimeInfo(id string) (string, error)
	RunJob(request spark.JobParams, sparkJobName string) (string, error)
}

func NewCmdSpark(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spark",
		Short: "Manage Spark jobs",
		Long:  "Manage Spark jobs",
		Example: heredoc.Doc(`
			$ snd spark request-status --id 54284cb9-8e58-4d92-93cb-6543
			$ snd spark runtime-info --id 54284cb9-8e58-4d92-93cb-6543
			$ snd spark logs --id 54284cb9-8e58-4d92-93cb-6543
			$ snd spark submit --job-name configuration-name --overrides ./overrides.json
			$ snd spark configuration --name configuration-name 
		`),
		GroupID: "spark",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&id, "id", "i", "", "Specify the  Job ID")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", beastURL, "Specify the service url")

	cmd.AddCommand(NewCmdSubmit(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRuntimeInfo(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRequestStatus(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdLogs(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdConfiguration(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdEncrypt(authServiceFactory, serviceFactory))

	return cmd
}
