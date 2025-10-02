package errors

import (
	"errors"
	"fmt"
)

var (
	ErrorInternalServer  = errors.New("internal server error")
	ErrFailedTracerSetup = errors.New("failed to setup new tracer")

	ErrorInvalidCredentials = errors.New("invalid credentials")
	ErrorPasswordEncryption = errors.New("failure to encrypt password, try a different one")
	ErrorDuplicatedAccount  = errors.New("there is already an account with that email address")

	ErrorSessionRetrieve     = errors.New("failure to retrieve session information")
	ErrorAccountRetrieve     = errors.New("failure to retrieve account information")
	ErrorAccountDataRetrieve = errors.New("failure to retrieve account data information")
	ErrorAccountDataUpdate   = errors.New("failure to update account data information")

	ErrorAccountCreation = errors.New("failure to create an account")

	ErrorFailOpenDBConnection = errors.New("fail to open DB Connection")
	ErrorFailedTracerSetup    = errors.New("failed to setup new tracer")

	ErrorExpiredToken = errors.New("authorization token expired")
	ErrorRenewedToken = errors.New("please renew auth token")
	ErrorInvalidToken = errors.New("authorization invalid")

	ErrorUnexpected     = errors.New("unexpected error")
	ErrorBadRequest     = errors.New("bad request")
	ErrorInvalidPGXRows = errors.New("invalid pgx rows mapping")

	ErrorMissingAuthHeader = errors.New("authorization header not provided")
	ErrorInvalidAuthHeader = errors.New("authorization header invalid")
	ErrorInvalidAuthType   = errors.New("authorization type invalid")

	// tests
	ErrorBCryptWrongLength = errors.New("bcrypt: password length exceeds 72 bytes")
	ErrorBCryptPassword    = errors.New("is not the hash of the given password")
	ErrorPGXMockDefault    = errors.New("pgx mock error")
	ErrorPGXTransFail      = errors.New("call to method BeginTx(), was not expected")

	ErrorSQLColumnReference = errors.New("column reference")
	ErrorSQLMissingField    = errors.New("cannot find field")
	ErrorSQLIncorrectArgs   = errors.New("incorrect argument number")
)

func New(caller string, err error) error {
	return fmt.Errorf("caller: %v, error: %v", caller, err)
}
