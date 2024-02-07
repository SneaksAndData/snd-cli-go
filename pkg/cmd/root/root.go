package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/auth"
	"snd-cli/pkg/cmd/claim"
	"snd-cli/pkg/cmd/ml"
	"snd-cli/pkg/cmd/spark"
	"snd-cli/pkg/cmd/spark/configuration"
	"snd-cli/pkg/cmd/spark/encrypt"
)

func NewCmdRoot() (*cobra.Command, error) {
	// Cmd represents the base command when called without any subcommands
	var cmd = &cobra.Command{
		Use:   "snd <command> <subcommand> [flags]",
		Short: "SnD CLI",
		Long:  `SnD CLI is a tool for interacting with various internal and external services in Sneaks & Data`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO: check that the user is authenticated before running most commands
			fmt.Println("To get started with SnD CLI, please run: snd login")
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
	cmd.AddCommand(auth.NewCmdAuth())
	cmd.AddCommand(claim.NewCmdClaim())
	cmd.AddCommand(ml.NewCmdAlgorithm())
	cmd.AddCommand(spark.NewCmdSpark())
	cmd.AddCommand(configuration.NewCmdConfiguration())
	cmd.AddCommand(encrypt.NewCmdEncrypt())
	return cmd, nil

}
