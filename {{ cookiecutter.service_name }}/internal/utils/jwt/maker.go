package jwt

import (
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
)

type Maker interface {
	CreateToken(account *models.Account, issued time.Time) (string, error)
	DecryptToken(token string) (*Session, error)
	VerifyToken(token string) (*Session, error)
	RenewToken(token string) (string, error)
}
