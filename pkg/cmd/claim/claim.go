// Package claim /*
package claim

import (
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util"
)

var env, authProvider, userId, claimProvider string

func NewCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "claim",
		Short:   "Manage claims",
		GroupID: "claim",
	}

	cmd.AddCommand(NewCmdUser())
	cmd.AddCommand(NewCmdAddClaim())
	cmd.AddCommand(NewCmdGetClaim())
	cmd.AddCommand(NewCmdRemoveClaim())

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&userId, "user", "u", "", "Specify the user ID")
	cmd.PersistentFlags().StringVarP(&claimProvider, "claims-provider", "p", "", "Specify the claim provider")

	return cmd
}

func InitClaimService(url string) (*claim.Service, error) {
	config := claim.Config{
		ClaimURL:     url,
		GetTokenFunc: util.ReadToken,
	}
	claimService, err := claim.New(config)
	if err != nil {
		log.Fatalf("Failed to create claim service: %v", err)
	}
	return claimService, nil
}
