package mock

import "github.com/influxdata/influxdb/cmd/influx/config"

// ConfigService mocks the ConfigService.
type ConfigService struct {
	WriteConfigsFn func(pp config.Configs) error
	ParseConfigsFn func() (config.Configs, error)
}

// WriteConfigs returns the write fn.
func (s *ConfigService) WriteConfigs(pp config.Configs) error {
	return s.WriteConfigsFn(pp)
}

// ParseConfigs returns the parse fn.
func (s *ConfigService) ParseConfigs() (config.Configs, error) {
	return s.ParseConfigsFn()
}
