package services

import (
	"github.com/SneaksAndData/esd-services-api-client-go/auth"
	"snd-cli/internal/config"
)

type ServiceProvider struct {
	cfg *config.ServiceConfig
}

func NewServiceProvider(cfg *config.ServiceConfig) *ServiceProvider {
	return &ServiceProvider{cfg: cfg}
}

func (sp *ServiceProvider) NewAuthService(env, provider string) (*auth.Service, error) {
	sp.cfg.AuthConfig.Env = env
	sp.cfg.AuthConfig.Provider = provider
	return auth.New(sp.cfg.AuthConfig)
}
