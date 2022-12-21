package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var ErrConfigFileNotFound = errors.New("config file not found")
var ErrConfigFileError = errors.New("read in config file error")

type Config struct {
	App          `mapstructure:"app" json:"app" yaml:"app"`
	DB           `mapstructure:"database" json:"database" yaml:"database"`
	FinnHub      `mapstructure:"finnHub" json:"finnHub" yaml:"finnHub"`
	AlphaVantage `mapstructure:"alphaVantage" json:"alphaVantage" yaml:"alphaVantage"`
	Log          `mapstructure:"log" json:"log" yaml:"log"`
}

func New() (error, *Config) {
	config := Config{}
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("etc/web_app")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ErrConfigFileNotFound, nil
		} else {
			return ErrConfigFileError, nil
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	return nil, &config
}
