package claim

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
	"strings"
)

func NewCmdGetClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   heredoc.Doc(`Retrieves claims assigned to an existing user`),
		Example: heredoc.Doc(`snd claim get -u user@ecco.com --claims-provider azuread`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("claim", env, url, authService)
			if err != nil {
				return err
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
		return "", fmt.Errorf("failed to retrieve claims for user %s for claim provider %s : %w", userId, claimProvider, err)
	}
	return response, nil
}
