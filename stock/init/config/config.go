package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App          `mapstructure:"app" json:"app" yaml:"app"`
	DB           `mapstructure:"database" json:"database" yaml:"database"`
	FinnHub      `mapstructure:"finnHub" json:"finnHub" yaml:"finnHub"`
	AlphaVantage `mapstructure:"alphaVantage" json:"alphaVantage" yaml:"alphaVantage"`
	Log          `mapstructure:"log" json:"log" yaml:"log"`
	Redis        `mapstructure:"redis" json:"redis" yaml:"redis"`
	Xueqiu       `mapstructure:"xueqiu" json:"xueqiu" yaml:"xueqiu"`
}

func New() (error, *Config) {
	config := Config{}
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err, nil
		} else {
			return err, nil
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	return nil, &config
}
