package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
)

var cr []string

func NewCmdRemoveClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a claim from an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return removeClaimRun(claimService)
		},
	}

	cmd.Flags().StringSliceVarP(&cr, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func removeClaimRun(claimService *claim.Service) error {
	// Add user claims
	response, err := claimService.RemoveClaim(userId, claimProvider, cr)
	if err != nil {
		log.Fatalf("Failed to remove claims for user %s with claim provider %s: %v", userId, claimProvider, err)
	}

	fmt.Println("Response:", response)
	return nil
}
