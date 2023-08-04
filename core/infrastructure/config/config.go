package config

type AppConfig struct {
	Debug  bool
	Db     DatabaseConfig
	Logger LoggerConfig
	Http   HttpConfig
}

type DatabaseConfig struct {
	Host          string `default:"localhost" validate:"required"`
	DbName        string `validate:"required"`
	Username      string `validate:"required"`
	Password      string `validate:"required"`
	Port          int    `default:"5432"`
	RetryAttempts uint   `default:"3"`
}

type LoggerConfig struct {
	Level     string
	Colorized bool
}

type HttpConfig struct {
	Port int `default:"2222"`
	EnableLogGin bool `default:"true"`
	GinMode string `default:"debug" validate:"eq=debug|eq=test|eq=release"`
}

type ConfigurationReader interface {
	LoadConfiguration(configPath string) error
}

type SimpleReader struct {
	conf AppConfig
}
