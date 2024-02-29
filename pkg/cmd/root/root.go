package root

import (
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/spf13/cobra"
	authCmd "snd-cli/pkg/cmd/auth"
	claimCmd "snd-cli/pkg/cmd/claim"
	mlCmd "snd-cli/pkg/cmd/ml"
	sparkCmd "snd-cli/pkg/cmd/spark"
	"snd-cli/pkg/cmdutil"
)

// AuthServiceFactory is a function type that creates a Service instance.
type AuthServiceFactory func(env, provider string) (*auth.Service, error)

type CacheToken func(token string) (string, error)

func NewCmdRoot() (*cobra.Command, error) {
	// Cmd represents the base command when called without any subcommands
	var cmd = &cobra.Command{
		Use:   "snd <service command group> <service command> [flags]",
		Short: "SnD CLI",
		Long:  `SnD CLI is a tool for interacting with various internal and external services in Sneaks & Data.`,
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

	authServiceFactory := cmdutil.NewAuthServiceFactory()
	serviceFactory := cmdutil.NewConcreteServiceFactory()

	cmd.SetVersionTemplate("Version: ")

	// Child commands
	cmd.AddCommand(authCmd.NewCmdAuth(authServiceFactory))
	cmd.AddCommand(claimCmd.NewCmdClaim(serviceFactory, authServiceFactory))
	cmd.AddCommand(mlCmd.NewCmdAlgorithm(serviceFactory, authServiceFactory))
	cmd.AddCommand(sparkCmd.NewCmdSpark(serviceFactory, authServiceFactory))
	return cmd, nil
}
