package config

type AppConfig struct {
	Debug bool
	Db    DatabaseConfig
}

type DatabaseConfig struct {
	Host          string `default:"localhost" validate:"required"`
	DbName        string `validate:"required"`
	Username      string `validate:"required"`
	Password      string `validate:"required"`
	Port          int    `default:"5432"`
	RetryAttempts uint    `default:"3"`
}

type ConfigurationReader interface {
	LoadConfiguration(configPath string) error
}

type SimpleReader struct {
	conf *AppConfig
}
