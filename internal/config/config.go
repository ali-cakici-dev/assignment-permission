package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string `env:"DATABASE_URI" envDefault:"mongodb://localhost:27017"`
	Port        int    `env:"PORT" envDefault:"8080"`
}

func LoadConfig(path, filename string) (*Config, error) {
	config := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, using environment variables")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	return &config, nil
}
