package errors

import (
	"errors"
)

var (
	ErrorInternalServer       = errors.New("internal server error")
	ErrorFailOpenDBConnection = errors.New("fail to open DB Connection")
	ErrorFailedTracerSetup    = errors.New("failed to setup new tracer")

	ErrorMissingAuthHeader = errors.New("authorization header not provided")
	ErrorInvalidAuthHeader = errors.New("authorization header invalid")
	ErrorInvalidAuthType   = errors.New("authorization type invalid")
)

func New(message string) error {
	return errors.New(message)
}
