package config

import "github.com/num30/config"

func (r *SimpleReader) LoadConfiguration(configPath string) error {
	if configPath == "" {
		configPath = ".env"
	}

	return config.NewConfReader(configPath).Read(r.conf)
}

func NewConfigProvider(configPath string) AppConfig {
	r := SimpleReader{
		conf: AppConfig{},
	}

	r.LoadConfiguration(configPath)
	return r.conf
}
