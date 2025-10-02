package authorization

import (
	"log"
	"strings"
	"testing"
	"time"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/enums"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/db/models"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/mocks"

	"github.com/stretchr/testify/require"
)

var (
	cfg = config.PasetoConfig{
		TOKEN_SYMETRIC_KEY:    mocks.RandomString(32),
		TOKEN_RENEW_DURATION:  time.Hour * 24,
		TOKEN_EXPIRE_DURATION: time.Hour * 24 * 5,
	}
	password = "12345678"

	account = models.Account{
		ID:        1,
		Tag:       mocks.RandomString(32),
		Firstname: mocks.RandomFirstname(),
		Lastname:  mocks.RandomLastname(),
		Password:  password,
		Email:     mocks.RandomEmail(),
		Role:      enums.User,
	}
	account2 = models.Account{
		ID:        2,
		Tag:       mocks.RandomString(32),
		Firstname: mocks.RandomFirstname(),
		Lastname:  mocks.RandomLastname(),
		Password:  password,
		Email:     mocks.RandomEmail(),
		Role:      enums.User,
	}
)

func TestAuthorizationMaker(t *testing.T) {
	maker, err := NewPasetoMaker(cfg)

	require.NoError(t, err)
	require.NotEmpty(t, maker)
}

func TestCreateAuthorizationToken(t *testing.T) {
	maker, _ := NewPasetoMaker(cfg)
	token, err := maker.CreateToken(&account, time.Now())

	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestVerifyAuthorizationTokenDecrypt(t *testing.T) {
	maker, _ := NewPasetoMaker(cfg)

	session, err := maker.VerifyToken(mocks.RandomString(50))
	require.Error(t, err)

	require.Contains(t, strings.ToLower(err.Error()), "authorization.verifytoken decrypt")
	require.Contains(t, strings.ToLower(err.Error()), "failed to decode token")
	require.Empty(t, session)
}

func TestVerifyAuthorizationExpiredToken(t *testing.T) {
	maker, _ := NewPasetoMaker(cfg)
	issuedAt := time.Now()
	token, _ := maker.CreateToken(&account, issuedAt)
	issuedAt = issuedAt.AddDate(0, 0, -2)
	expired_token, _ := maker.CreateToken(&account2, issuedAt)

	session, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, session.ID, account.ID)

	expired_session, err := maker.VerifyToken(expired_token)
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "authorization.verifytoken invalid")
	require.Empty(t, expired_session)
}

func TestVerifyAuthorizationRenewToken(t *testing.T) {
	maker, _ := NewPasetoMaker(cfg)
	issuedAt := time.Now()
	token, _ := maker.CreateToken(&account, issuedAt)
	session, _ := maker.VerifyToken(token)

	renewed_token, err := maker.RenewToken(token)
	require.NoError(t, err)
	renewed_session, _ := maker.VerifyToken(renewed_token)

	require.Equal(t, renewed_session.IssuedAt, session.IssuedAt)
	require.Equal(t, renewed_session.RenewAt, session.RenewAt)
	require.Equal(t, renewed_session.ExpiredAt, session.ExpiredAt)

	issuedAt = issuedAt.AddDate(0, 0, -1)
	expired_token, _ := maker.CreateToken(&account, issuedAt)
	session, _ = maker.DecryptToken(expired_token)
	renewed_token, _ = maker.RenewToken(expired_token)
	renewed_session, _ = maker.VerifyToken(renewed_token)

	log.Printf("initial %s", session.RenewAt)
	log.Printf("renewed %s", renewed_session.RenewAt)

	require.Equal(t, renewed_session.IssuedAt, session.IssuedAt)
	require.Greater(t, renewed_session.RenewAt, session.RenewAt)
	require.Equal(t, renewed_session.ExpiredAt, session.ExpiredAt)
}
