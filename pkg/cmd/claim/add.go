package claim

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var claimsToAdd []string

func NewCmdAddClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Short:   heredoc.Doc(`Add a new claim to an existing user`),
		Example: heredoc.Doc(`snd claim add -c "service.test.sneaksanddata.com/.*:.*" -u user@ecco.com --claims-provider azuread`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
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
			resp, err := addClaimRun(service.(*claim.Service), userId, claimProvider, claimsToAdd)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringSliceVarP(&claimsToAdd, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func addClaimRun(claimService Service, userId, claimProvider string, claimsToAdd []string) (string, error) {
	// Validate claims
	for _, c := range claimsToAdd {
		if !util.ValidateClaim(c) {
			return "", fmt.Errorf("invalid claim format: Ensure the claim string follows the pattern 'path:method'. Please review your claim string: %s", c)
		}
	}
	// Add user claims
	response, err := claimService.AddClaim(userId, claimProvider, claimsToAdd)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to add claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to add claims for user %s with claim provider %s: %w", userId, claimProvider, err)
	}
	prettifyResponse, err := util.PrettifyJSON(response)
	if err != nil {
		return "", fmt.Errorf("failed to prettify response: %w", err)
	}
	return prettifyResponse, nil
}
