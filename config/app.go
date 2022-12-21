package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	AppName string `mapstructure:"appname" json:"appname" yaml:"appname"`
}
