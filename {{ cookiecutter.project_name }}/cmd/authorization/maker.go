package authorization

import (
	"time"
)

// maker interface for managing tokens
type Maker interface {
	// creates a new token for specific username and duration
	CreateToken(id int, email string, duration time.Duration) (string, error)

	// verify if the token is valid
	VerifyToken(token string) (*Payload, error)
}
