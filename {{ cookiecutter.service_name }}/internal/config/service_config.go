package config

import "{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts"

type ServiceConfig struct {
	NAME    string                `mapstructure:"name"`
	PORT    string                `mapstructure:"port"`
	VERSION string                `mapstructure:"version"`
	MODE    consts.RunMode        `mapstructure:"mode"`
	ENV     consts.RunEnvironment `mapstructure:"environment"`
}
