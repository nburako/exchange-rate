package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mongodb     MongodbConfig
	ExternalApi ExternalApiConfig
}

type MongodbConfig struct {
	Db         string
	Database   string
	Collection string
}

type ExternalApiConfig struct {
	Endpoint string
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := loadConfig(configPath)

	if err != nil {
		return nil, err
	}

	config, err := parseConfig(cfgFile)

	if err != nil {
		return nil, err
	}

	return config, err
}

func loadConfig(configPath string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(configPath)
	v.AddConfigPath("../")
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}

		return nil, err
	}

	return v, err
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode, %v", err)

		return nil, err
	}

	return &c, nil
}
