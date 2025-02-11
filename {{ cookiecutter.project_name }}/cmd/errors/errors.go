package errors

import (
	"errors"
)

var (
	ErrorInternalServer = errors.New("internal server error")

	ErrorInvalidCredentials = errors.New("invalid credentials")
	ErrorPasswordEncryption = errors.New("failure to encrypt password, try a different one")
	ErrorDuplicatedAccount  = errors.New("there is already an account with that email address")
	ErrorAccountRetrieve    = errors.New("failure to retrieve account information")
	ErrorAccountCreation    = errors.New("failure to create an account")

	ErrorFailOpenDBConnection = errors.New("fail to open DB Connection")
	ErrorFailedTracerSetup    = errors.New("failed to setup new tracer")

	ErrorInvalidAuth = errors.New("authorization invalid")
	ErrorExpiredAuth = errors.New("authorization expired")

	ErrorMissingAuthHeader = errors.New("authorization header not provided")
	ErrorInvalidAuthHeader = errors.New("authorization header invalid")
	ErrorInvalidAuthType   = errors.New("authorization type invalid")
)

func New(message string) error {
	return errors.New(message)
}
