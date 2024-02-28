package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmdutil"
	"strings"
)

func NewCmdGetClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves claims assigned to an existing user",
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
			resp, err := getClaimRun(service.(*claim.Service), userId, claimProvider)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	return cmd
}

func getClaimRun(claimService Service, userId, claimProvider string) (string, error) {
	// Retrieve user claims
	response, err := claimService.GetClaim(userId, claimProvider)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, err)
	}
	return response, nil
}
