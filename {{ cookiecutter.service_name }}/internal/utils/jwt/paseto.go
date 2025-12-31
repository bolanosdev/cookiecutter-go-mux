package jwt

import (
	"fmt"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"

	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	config config.PasetoConfig
}

func NewPasetoMaker(cfg config.PasetoConfig) (Maker, error) {
	if len(cfg.TOKEN_SYMETRIC_KEY) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		config: cfg,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(account *models.Account, issued time.Time) (string, error) {
	renewAt := maker.config.TOKEN_RENEW_DURATION
	expireAt := maker.config.TOKEN_EXPIRE_DURATION

	session := NewSession(account, issued, renewAt, expireAt)
	return maker.paseto.Encrypt([]byte(maker.config.TOKEN_SYMETRIC_KEY), session, nil)
}

func (maker *PasetoMaker) DecryptToken(token string) (*Session, error) {
	session := &Session{}

	err := maker.paseto.Decrypt(token, []byte(maker.config.TOKEN_SYMETRIC_KEY), session, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Authorization.DecryptToken decrypt")
	}

	return session, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Session, error) {
	session, err := maker.DecryptToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "Authorization.VerifyToken decrypt")
	}

	validated_session, err := session.Valid()
	if err != nil {
		return nil, errors.Wrap(err, "Authorization.VerifyToken invalid")
	}

	return validated_session, nil
}

func (maker *PasetoMaker) RenewToken(token string) (string, error) {
	session := &Session{}

	err := maker.paseto.Decrypt(token, []byte(maker.config.TOKEN_SYMETRIC_KEY), session, nil)
	if err != nil {
		return "", errors.Wrap(err, "Authorization.RenewToken decrypt")
	}

	renewed_session, err := session.Renew(maker.config.TOKEN_RENEW_DURATION)
	if err != nil {
		return "", errors.Wrap(err, "Authorization.RenewToken renew")
	}

	return maker.paseto.Encrypt([]byte(maker.config.TOKEN_SYMETRIC_KEY), renewed_session, nil)
}
