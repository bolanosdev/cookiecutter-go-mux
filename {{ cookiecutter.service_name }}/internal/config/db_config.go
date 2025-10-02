package config

type DBConfig struct {
	HOSTNAME string `mapstructure:"hostname"`
	PORT     string `mapstructure:"port"`
	USERNAME string `mapstructure:"username"`
	PASSWORD string `mapstructure:"password"`
	DATABASE string `mapstructure:"database"`
	SSL      string `mapstructure:"ssl"`
}
