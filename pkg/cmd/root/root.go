package root

import (
	apiAuth "github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/spf13/cobra"
	authCmd "snd-cli/pkg/cmd/auth"
	claimCmd "snd-cli/pkg/cmd/claim"
	mlCmd "snd-cli/pkg/cmd/ml"
	sparkCmd "snd-cli/pkg/cmd/spark"
)

// AuthServiceFactory is a function type that creates a Service instance.
type AuthServiceFactory func(env, provider string) (*apiAuth.Service, error)

type CacheToken func(token string) (string, error)

func NewCmdRoot() (*cobra.Command, error) {
	// Cmd represents the base command when called without any subcommands
	var cmd = &cobra.Command{
		Use:   "snd <service command group> <service command> [flags]",
		Short: "SnD CLI",
		Long:  `SnD CLI is a tool for interacting with various internal and external services in Sneaks & Data.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Check that the user is authenticated before running most commands
			// fmt.Println("To get started with SnD CLI, please run: snd login")
			return nil
		},
	}
	cmd.AddGroup(&cobra.Group{
		ID:    "auth",
		Title: "Auth Commands",
	})

	cmd.AddGroup(&cobra.Group{
		ID:    "claim",
		Title: "Claim Commands",
	})

	cmd.AddGroup(&cobra.Group{
		ID:    "ml",
		Title: "ML Algorithm Commands",
	})

	cmd.AddGroup(&cobra.Group{
		ID:    "spark",
		Title: "Spark Commands",
	})

	// Child commands
	cmd.AddCommand(authCmd.NewCmdAuth())
	cmd.AddCommand(claimCmd.NewCmdClaim())
	cmd.AddCommand(mlCmd.NewCmdAlgorithm())
	cmd.AddCommand(sparkCmd.NewCmdSpark())
	return cmd, nil
}
