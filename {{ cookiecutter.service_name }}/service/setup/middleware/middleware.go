package middleware

import (
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/setup/authorization"
)

type Middleware struct {
	Config config.AppConfig
	Paseto authorization.Maker
}

func NewMiddleware(cfg config.AppConfig, paseto authorization.Maker) Middleware {
	return Middleware{
		Config: cfg,
		Paseto: paseto,
	}
}
