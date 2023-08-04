package config

import "github.com/num30/config"

func (r *SimpleReader) LoadConfiguration(configPath string) error {
	if configPath == "" {
		configPath = ".env"
	}

	return config.NewConfReader(configPath).Read(r.conf)
}

func NewReader(configPath string) SimpleReader {
	return SimpleReader{
		conf: &AppConfig{},
	}
}
