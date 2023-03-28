package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PoolSize int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
}
