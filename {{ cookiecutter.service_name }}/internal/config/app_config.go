package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	ConfigMgr struct {
		path string
	}
	AppConfig struct {
		SERVICE       ServiceConfig `mapstructure:"service"`
		DATABASE      DBConfig      `mapstructure:"db"`
		PASETO        PasetoConfig  `mapstructure:"paseto"`
		OBSERVABILITY ObsConfig     `mapstructure:"observability"`
	}
)

func NewConfigMgr(path string) ConfigMgr {
	return ConfigMgr{
		path: path,
	}
}

func (cfg ConfigMgr) Load() (config AppConfig) {
	viper.AddConfigPath(cfg.path)
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
