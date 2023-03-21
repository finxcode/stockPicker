package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var ErrConfigFileNotFound = errors.New("config file not found")
var ErrConfigFileError = errors.New("read in config file error")

type Config struct {
	FinnHub `mapstructure:"finnHub" json:"finnHub" yaml:"finnHub"`
}

type FinnHub struct {
	BaseUrl string `mapstructure:"baseUrl" json:"baseUrl" yaml:"baseUrl"`
	Stock   string `mapstructure:"stock" json:"stock" yaml:"stock"`
	Symbol  string `mapstructure:"symbol" json:"symbol" yaml:"symbol"`
	ApiKey  string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`
}

func New() (error, *Config) {
	config := Config{}
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("./ext/finnhub/")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ErrConfigFileNotFound, nil
		} else {
			return err, nil
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	return nil, &config
}
