package auth

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmd/util"
)

var env, provider string

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Get internal authorization token",
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginRun()
		},
	}

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")

	return cmd
}

func loginRun() error {
	fmt.Println("Login")
	tokenURL := fmt.Sprintf(boxerBaseURL, env)
	config := auth.Config{
		TokenURL: tokenURL,
		Env:      env,
		Provider: provider,
	}

	// Create a new instance of the auth service
	authService, err := auth.New(config)
	if err != nil {
		log.Fatalf("Failed to create auth service: %v", err)
	}
	// Retrieve token
	token, err := authService.GetBoxerToken()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	cachedToken, err := util.CacheToken(token)
	if err != nil {
		log.Fatalf("Failed to cache token: %v", err)
	}
	fmt.Println("Token:", cachedToken)
	return nil
}
