package util

import (
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	DBSource     string `mapstructure:"DB_SOURCE"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	env := os.Getenv("GO_ENV")
	envFile := "app.env"
	if env == "" || env == "local" {
		envFile += ".local"
	}
	viper.SetConfigFile(path + "/" + envFile)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
