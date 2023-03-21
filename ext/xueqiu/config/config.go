package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var ErrConfigFileNotFound = errors.New("config file not found")
var ErrConfigFileError = errors.New("read in config file error")

type Config struct {
	AlphaVantage `mapstructure:"alphaVantage" json:"alphaVantage" yaml:"alphaVantage"`
}

type AlphaVantage struct {
	ApiKey string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`
}

func New() (*Config, error) {
	config := Config{}
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("./ext/finnhub/")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, ErrConfigFileNotFound
		} else {
			return nil, ErrConfigFileError
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	return &config, nil
}
