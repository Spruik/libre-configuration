package port

import (
	"github.com/Spruik/libre-configuration/shared"
)

type ConfigProviderIF interface {
	Initialize(filePath string)
	GetComponentConfig(component string) *shared.ConfigItem
	GetConfigEntry(component string, key string) (string, error)
	GetConfigEntryWithDefault(component string, key string, dflt string) (string, error)
	GetConfigStanza(component string, top string) (*shared.ConfigItem, error)
}
