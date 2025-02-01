package cfg

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	SERVICE_NAME  string       `mapstructure:"service_name"`
	DATABASE      DBConfig     `mapstructure:"db"`
	PASETO        PasetoConfig `mapstructure:"paseto"`
	OBSERVABILITY ObsConfig    `mapstructure:"observability"`
}

func Load(path string) (config AppConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failing to read viper config %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("failing to Unmarshal viper config %v", err)
	}

	return
}
