package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost          string `mapstructure:"DB_HOST"`
	DBName          string `mapstructure:"DB_NAME"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	AccessTokenSrc  string `mapstructure:"ACCESS_TOKEN_SRC "`
	RefreshTokenSrc string `mapstructure:"REFRESH_TOKEN_SRC"`
	ClientID        string `mapstructure:"CLIENT_ID"`
	ClientSecret    string `mapstructure:"CLIENT_SECRET"`
	RedirectURL     string `mapstructure:"REDIRECT_URL"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "ACCESS_TOKEN_SRC", "REFRESH_TOKEN_SRC",
	"CLIENT_ID", "CLIENT_SECRET", "REDIRECT_URL",
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
