package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	AppName string `mapstructure:"appName" json:"appName" yaml:"appName"`
}
