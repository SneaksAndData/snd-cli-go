package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func NewCmdGetClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves claims assigned to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return getClaimRun(claimService)
		},
	}

	return cmd
}

func getClaimRun(claimService *claim.Service) error {
	// Retrieve user claims
	response, err := claimService.GetClaim(userId, claimProvider)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			log.Fatalf("Failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		log.Fatalf("Failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, err)
	}
	fmt.Println(response)
	return nil
}
