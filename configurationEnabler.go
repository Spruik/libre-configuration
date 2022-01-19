package libreConfig

import (
	"errors"

	"github.com/Spruik/libre-configuration/shared"
)

const errMessageNotInitalized = "configurationEnabler not initialized. Please call `Initialize(\"path/to/file.json\")`"

type ConfigurationEnabler struct {
	category string
}

func (s *ConfigurationEnabler) SetConfigCategory(cat string) {
	if s == nil {
		panic("cannot set category of uninitialized")
	}
	s.category = cat
}
func (s *ConfigurationEnabler) GetConfigItem(key string) (string, error) {
	if cfgSvc == nil {
		return "", errors.New(errMessageNotInitalized)
	}
	return cfgSvc.GetConfigEntry(s.category, key)
}
func (s *ConfigurationEnabler) GetConfigItemWithDefault(key string, dflt string) (string, error) {
	if cfgSvc == nil {
		return dflt, errors.New(errMessageNotInitalized)
	}
	return cfgSvc.GetConfigEntryWithDefault(s.category, key, dflt)
}
func (s *ConfigurationEnabler) GetConfigStanza(key string) (*shared.ConfigItem, error) {
	if cfgSvc == nil {
		return nil, errors.New(errMessageNotInitalized)
	}
	return cfgSvc.GetConfigStanza(s.category, key)
}
