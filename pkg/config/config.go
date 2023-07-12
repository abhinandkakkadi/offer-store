package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DB_USERNAME     string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD     string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT  string `mapstructure:"DB_PORT"`
}

var envs = []string{
	"DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_HOST", "DB_PORT",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil

}
