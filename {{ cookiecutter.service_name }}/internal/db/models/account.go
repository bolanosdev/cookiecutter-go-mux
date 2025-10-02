package models

import "time"

type Account struct {
	ID        int
	Email     string
	Password  string
	Role      Role
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
