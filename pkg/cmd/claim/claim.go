// Package claim /*
package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util/token"
)

const boxerClaimBaseURL = "https://boxer-claim.%s.sneaksanddata.com"

var env, authProvider, userId, claimProvider string

type Service interface {
	AddClaim(user string, provider string, claims []string) (string, error)
	GetClaim(user string, provider string) (string, error)
	RemoveClaim(user string, provider string, claims []string) (string, error)
	AddUser(user string, provider string) (string, error)
	RemoveUser(user string, provider string) (string, error)
}

func NewCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "claim",
		Short:   "Manage claims",
		GroupID: "claim",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&userId, "user", "u", "", "Specify the user ID")
	cmd.PersistentFlags().StringVarP(&claimProvider, "claims-provider", "p", "", "Specify the claim provider")

	var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
	if err != nil {
		log.Fatalf("Failed to initialize claim service: %v", err)
	}
	cmd.AddCommand(NewCmdUser(claimService))
	cmd.AddCommand(NewCmdAddClaim(claimService))
	cmd.AddCommand(NewCmdGetClaim(claimService))
	cmd.AddCommand(NewCmdRemoveClaim(claimService))

	return cmd
}

func InitClaimService(url string) (*claim.Service, error) {
	tc := token.TokenCache{}
	config := claim.Config{
		ClaimURL:     url,
		GetTokenFunc: tc.ReadToken,
	}
	claimService, err := claim.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create claim service: %v", err)
	}
	return claimService, nil
}
