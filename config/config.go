package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

func ReadConfigFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("reading configuration: %w", err)
	}

	cfg := Config{
		Port:       viper.GetInt("port"),
		DbHost:     viper.GetString("db_host"),
		DbPort:     viper.GetInt("db_port"),
		DbUser:     viper.GetString("db_user"),
		DbPassword: viper.GetString("db_password"),
		DbName:     viper.GetString("db_name"),
	}

	return cfg, nil
}
