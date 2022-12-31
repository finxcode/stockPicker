package config

type FinnHub struct {
	BaseUrl string `mapstructure:"baseUrl" json:"baseUrl" yaml:"baseUrl"`
	Stock   string `mapstructure:"stock" json:"stock" yaml:"stock"`
	Symbol  string `mapstructure:"symbol" json:"symbol" yaml:"symbol"`
	ApiKey  string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`
}
