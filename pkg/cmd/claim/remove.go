package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var cr []string

func NewCmdRemoveClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a claim from an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("claim", env, authService)
			if err != nil {
				return err
			}
			resp, err := removeClaimRun(service.(*claim.Service), userId, claimProvider, cr)
			if err == nil {
				log.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringSliceVarP(&cr, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func removeClaimRun(claimService Service, userId, claimProvider string, cr []string) (string, error) {
	// Add user claims
	response, err := claimService.RemoveClaim(userId, claimProvider, cr)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to remove claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to remove claims for user %s with claim provider %s: %w", userId, claimProvider, err)
	}

	return response, nil
}
