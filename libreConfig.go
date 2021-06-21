package libreConfig

import (
	"github.com/Spruik/libre-configuration/internal/core/service"
	"github.com/Spruik/libre-configuration/internal/implementation"
)

var cfgSvc *service.ConfigProviderService

func Initialize(configFilePath string) {
	cfgSvc = service.NewConfigProviderService(implementation.NewConfigProviderFile())
	cfgSvc.Initialize(configFilePath)
}

func GetConfigService() *service.ConfigProviderService {
	return cfgSvc
}
