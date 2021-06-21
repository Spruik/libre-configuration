package libreConfig

import "github.com/Spruik/libre-configuration/shared"

type ConfigurationEnabler struct {
	category string
}

func (s *ConfigurationEnabler) SetConfigCategory(cat string) {
	s.category = cat
}
func (s *ConfigurationEnabler) GetConfigItem(key string) (string, error) {
	return cfgSvc.GetConfigEntry(s.category, key)
}
func (s *ConfigurationEnabler) GetConfigItemWithDefault(key string, dflt string) (string, error) {
	return cfgSvc.GetConfigEntryWithDefault(s.category, key, dflt)
}
func (s *ConfigurationEnabler) GetConfigStanza(key string) (*shared.ConfigItem, error) {
	return cfgSvc.GetConfigStanza(s.category, key)
}
