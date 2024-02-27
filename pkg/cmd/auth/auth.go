package auth

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util/token"
)

var env, provider string

const boxerBaseURL = "https://boxer.%s.sneaksanddata.com"

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Get internal authorization token",
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var authService, err = InitAuthService()
			if err != nil {
				log.Fatalf("Failed to initialize auth service: %v", err)
			}
			return loginRun(authService)
		},
	}

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")

	return cmd
}

func InitAuthService() (*auth.Service, error) {
	config := auth.Config{
		TokenURL: fmt.Sprintf(boxerBaseURL, env),
		Env:      env,
		Provider: provider,
	}
	authService, err := auth.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth service: %v", err)
	}
	return authService, nil
}

func loginRun(authService token.AuthService) error {
	tokenProvider := token.NewProvider(authService)
	cachedToken, err := tokenProvider.GetToken() // Fetch and cache the token.
	if err != nil {
		log.Fatalf("Error logging: %v", err)
	}
	log.Println("Login successful.")

	log.Println("Token:", cachedToken)
	return nil
}
