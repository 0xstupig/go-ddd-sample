package config

import (
	"fmt"
	"github.com/num30/config"
)

func (r *SimpleReader) LoadConfiguration(configPath string) error {
	if configPath == "" {
		configPath = "env"
	}

	reader := config.NewConfReader(configPath).WithSearchDirs(".")
	return reader.Read(&r.conf)
}

func NewConfigProvider(configPath string) AppConfig {
	r := SimpleReader{
		conf: AppConfig{},
	}

	err := r.LoadConfiguration(configPath)
	if err != nil {
		panic(fmt.Errorf("load config failed: %v \n", err))
	}
	fmt.Printf("conf: %+v \n", r.conf)
	return r.conf
}
