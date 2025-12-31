package jwt

import (
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"

	"github.com/pkg/errors"
)

type Session struct {
	IssuedAt  time.Time `json:"issued_at"`
	RenewAt   time.Time `json:"rewnew_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ID        int       `json:"id"`
}

func NewSession(
	account *models.Account,
	issued time.Time,
	renewAt time.Duration,
	expireAt time.Duration,
) *Session {
	payload := &Session{
		ID:        account.ID,
		Email:     account.Email,
		IssuedAt:  issued,
		RenewAt:   issued.Add(renewAt),
		ExpiredAt: issued.Add(expireAt),
	}

	return payload
}

func (payload *Session) Valid() (*Session, error) {
	if time.Now().After(payload.RenewAt) {
		return nil, errors.Wrap(errors.New("authorization.payload expired"), "session has exceeded renew time")
	}

	return payload, nil
}

func (payload *Session) Renew(renewDuration time.Duration) (*Session, error) {
	if time.Now().After(payload.ExpiredAt) {
		return nil, errors.Wrap(errors.New("authorization.payload expired"), "session has exceeded expire time")
	}

	if time.Now().After(payload.RenewAt) {
		payload := Session{
			ID:        payload.ID,
			Email:     payload.Email,
			IssuedAt:  payload.IssuedAt,
			RenewAt:   payload.RenewAt.Add(renewDuration),
			ExpiredAt: payload.ExpiredAt,
		}
		return &payload, nil
	}

	return payload, nil
}
