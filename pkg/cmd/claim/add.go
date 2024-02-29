package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var ca []string

func NewCmdAddClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new claim to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("claim", env, authService)
			if err != nil {
				log.Fatal(err)
			}
			if err != nil {
				return err
			}
			resp, err := addClaimRun(service.(*claim.Service), userId, claimProvider, ca)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringSliceVarP(&ca, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	cmd.MarkFlagRequired("claims")
	return cmd
}

func addClaimRun(claimService Service, userId, claimProvider string, ca []string) (string, error) {
	// Add user claims
	response, err := claimService.AddClaim(userId, claimProvider, ca)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to add claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to add claims for user %s with claim provider %s: %w", userId, claimProvider, err)
	}
	return response, nil
}
