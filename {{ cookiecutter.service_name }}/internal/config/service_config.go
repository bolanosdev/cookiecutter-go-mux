package config

import "{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts/enums"

type ServiceConfig struct {
	NAME    string               `mapstructure:"name"`
	PORT    string               `mapstructure:"port"`
	VERSION string               `mapstructure:"version"`
	MODE    enums.RunMode        `mapstructure:"mode"`
	ENV     enums.RunEnvironment `mapstructure:"environment"`
}
