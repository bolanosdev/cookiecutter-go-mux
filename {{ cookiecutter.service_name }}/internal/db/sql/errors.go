package sql

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

// QueryError wraps an error with metadata that can be extracted for observability
// without exposing sensitive data in error messages
type QueryError struct {
	Err      error
	Args     pgx.StrictNamedArgs
	Location string
}

func (qe *QueryError) Error() string {
	return fmt.Sprintf("%s: %v", qe.Location, qe.Err)
}

func (qe *QueryError) Unwrap() error {
	return qe.Err
}

// Sensitive keywords that should be masked in error metadata
var sensitiveKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"secret",
	"token",
	"api_key",
	"apikey",
	"access_token",
	"refresh_token",
	"auth",
	"authorization",
	"credit_card",
	"card_number",
	"cvv",
	"ssn",
	"social_security",
}

// maskSensitiveArgs creates a copy of args with sensitive values masked
func maskSensitiveArgs(args pgx.StrictNamedArgs) map[string]interface{} {
	masked := make(map[string]interface{})

	for key, value := range args {
		// Check if the key contains any sensitive keyword
		isSensitive := false
		lowerKey := strings.ToLower(key)

		for _, keyword := range sensitiveKeywords {
			if strings.Contains(lowerKey, keyword) {
				isSensitive = true
				break
			}
		}

		if isSensitive {
			masked[key] = "***"
		} else {
			masked[key] = value
		}
	}

	return masked
}

// GetMetadata returns the metadata for observability tools with sensitive data masked
func (qe *QueryError) GetMetadata() map[string]interface{} {
	maskedArgs := maskSensitiveArgs(qe.Args)

	return map[string]interface{}{
		"location": qe.Location,
		"args":     maskedArgs,
	}
}

// NewQueryError creates a new QueryError with metadata
// Args parameter is optional - pass nil or omit if no args
func NewQueryError(err error, location string, args *pgx.StrictNamedArgs) error {
	var queryArgs pgx.StrictNamedArgs
	if args != nil {
		queryArgs = maskSensitiveArgs(*args)
	}

	return &QueryError{
		Err:      err,
		Args:     queryArgs,
		Location: location,
	}
}
