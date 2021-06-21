package service

import (
	"github.com/Spruik/libre-configuration/internal/core/port"
	"github.com/Spruik/libre-configuration/shared"
)

type ConfigProviderService struct {
	port port.ConfigProviderIF
}

func NewConfigProviderService(port port.ConfigProviderIF) *ConfigProviderService {
	return &ConfigProviderService{
		port: port,
	}
}

func (s *ConfigProviderService) Initialize(filePath string) {
	s.port.Initialize(filePath)
}
func (s *ConfigProviderService) GetConfigEntry(component string, key string) (string, error) {
	return s.port.GetConfigEntry(component, key)
}
func (s *ConfigProviderService) GetConfigStanza(component string, top string) (*shared.ConfigItem, error) {
	return s.port.GetConfigStanza(component, top)
}
func (s *ConfigProviderService) GetComponentConfig(component string) *shared.ConfigItem {
	return s.port.GetComponentConfig(component)
}
func (s *ConfigProviderService) GetConfigEntryWithDefault(component string, key string, dflt string) (string, error) {
	return s.port.GetConfigEntryWithDefault(component, key, dflt)
}
