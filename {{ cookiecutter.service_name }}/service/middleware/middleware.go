package middleware

import (
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/jwt"
)

type Middleware struct {
	Config config.AppConfig
	Paseto jwt.Maker
}

func NewMiddleware(cfg config.AppConfig, paseto jwt.Maker) Middleware {
	return Middleware{
		Config: cfg,
		Paseto: paseto,
	}
}
