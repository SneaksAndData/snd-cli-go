package auth

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/token"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var env, provider string

func NewCmdAuth(authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Get internal authorization token",
		Long:  "Retrieve the internal authorization token generated by Boxer by providing an authentication provider",
		Example: heredoc.Doc(`
			$ snd login -a azuread -e test
			$ snd login --auth-provider azuread --env production
		`),
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, provider)
			if err != nil {
				return fmt.Errorf("failed to initialize auth service: %w", err)
			}
			return loginRun(authService)
		},
	}

	helpStr := []string{
		"Specify the OAuth provider name ",
		"For in-cluster Kubernetes auth specify name of your kubernetes cluster context prefixed with `k8s`",
		"for example `k8s-esd-airflow-dev-0`",
	}

	cmd.PersistentFlags().StringVarP(&env, "env", "e", cmdutil.BaseEnvironment, "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", strings.Join(helpStr, "\n"))

	return cmd
}

func loginRun(authService token.AuthService) error {
	tokenProvider, err := token.NewProvider(authService, env)
	if err != nil {
		return fmt.Errorf("unable to create token provider: %w", err)
	}
	cachedToken, err := tokenProvider.GetToken() // Fetch and cache the token.
	if err != nil {
		return fmt.Errorf("unable to get the token: %w", err)
	}
	fmt.Println("Login successful.")
	fmt.Println("Token:", cachedToken)
	return nil
}
