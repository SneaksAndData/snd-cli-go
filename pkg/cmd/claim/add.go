package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var ca []string

func NewCmdAddClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new claim to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {

			var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return addClaimRun(claimService)
		},
	}
	cmd.Flags().StringSliceVarP(&ca, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func addClaimRun(claimService *claim.Service) error {
	// Add user claims
	response, err := claimService.AddClaim(userId, claimProvider, ca)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			log.Fatalf("Failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		log.Fatalf("Failed to add claims for user %s with claim provider %s: %v", userId, claimProvider, err)
	}

	fmt.Println("Response:", response)
	return nil
}
