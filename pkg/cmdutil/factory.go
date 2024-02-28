package cmdutil

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"snd-cli/pkg/cmd/util/token"
)

const boxerURL = "https://boxer.%s.sneaksanddata.com"
const boxerClaimURL = "https://boxer-claim.%s.sneaksanddata.com"
const crystalURL = "https://crystal.%s.sneaksanddata.com"
const beastURL = "https://beast-v3.%s.sneaksanddata.com"

// AuthServiceFactory is responsible for creating instances of AuthService.
// It encapsulates the logic required to configure and instantiate an AuthService.
type AuthServiceFactory struct{}

// NewAuthServiceFactory creates a new instance of AuthServiceFactory.
func NewAuthServiceFactory() *AuthServiceFactory {
	return &AuthServiceFactory{}
}

// CreateAuthService creates and returns an instance of AuthService based on the provided
// environment (env) and provider. This method configures the AuthService with environment-
// specific settings, such as the TokenURL, ensuring the AuthService is tailored to operate
// within the specified environment.
func (f *AuthServiceFactory) CreateAuthService(env, provider string) (*auth.Service, error) {
	config := auth.Config{
		TokenURL: fmt.Sprintf(boxerURL, env),
		Env:      env,
		Provider: provider,
	}
	authService, err := auth.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth service: %v", err)
	}
	return authService, nil
}

// ServiceFactory defines an interface for factories capable of creating different types of services.
type ServiceFactory interface {
	CreateService(serviceType string, env string, authService token.AuthService) (interface{}, error)
}

// ConcreteServiceFactory implements the ServiceFactory interface, providing concrete logic to create specific
// service instances.
type ConcreteServiceFactory struct{}

// NewConcreteServiceFactory initializes a new instance of ConcreteServiceFactory.
// It requires no parameters, offering a simple way to obtain a factory capable of creating various services.
func NewConcreteServiceFactory() *ConcreteServiceFactory {
	return &ConcreteServiceFactory{}
}

// CreateService dynamically creates and returns a service instance based on the specified serviceType,
// environment, and an instance of AuthService for authentication purposes. This method supports creating
// multiple types of services, each configured according to the environment and authentication requirements.
//
// Parameters:
//
//	serviceType: A string identifier for the type of service to create (e.g., "claim", "algorithm", "spark").
//	env: The environment in which the service will operate (e.g., "test", "production").
//	authService: An instance of AuthService used for authenticating the service's requests.
//
// Returns:
//
//	An interface{} representing the created service, which should be type-asserted to the appropriate service type.
//	An error if the service creation fails or if an unknown service type is specified.
func (f *ConcreteServiceFactory) CreateService(serviceType, env string, authService token.AuthService) (interface{}, error) {
	switch serviceType {
	case "claim":
		return initClaimService(env, authService)
	case "algorithm":
		return initAlgorithmService(env, authService)
	case "spark":
		return initSparkService(env, authService)
	default:
		return nil, fmt.Errorf("unknown service type: %s", serviceType)
	}
}

func initClaimService(env string, authService token.AuthService) (*claim.Service, error) {
	tp := token.NewProvider(authService)
	config := claim.Config{
		ClaimURL:     fmt.Sprintf(boxerClaimURL, env),
		GetTokenFunc: tp.GetToken,
	}
	claimService, err := claim.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create claim service: %v", err)
	}
	return claimService, nil
}

func initAlgorithmService(env string, authService token.AuthService) (*algorithm.Service, error) {
	tp := token.NewProvider(authService)
	config := algorithm.Config{
		SchedulerURL: fmt.Sprintf(crystalURL, env),
		APIVersion:   "v1.2",
		GetTokenFunc: tp.GetToken,
	}

	algorithmService, err := algorithm.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create algorithm service: %v", err)
	}
	return algorithmService, nil
}

func initSparkService(env string, authService token.AuthService) (*spark.Service, error) {
	tp := token.NewProvider(authService)
	config := spark.Config{
		BaseURL:      fmt.Sprintf(beastURL, env),
		GetTokenFunc: tp.GetToken,
	}

	sparkService, err := spark.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create spark service: %v", err)
	}
	return sparkService, nil

}
