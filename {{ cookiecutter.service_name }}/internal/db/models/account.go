package models

import "time"

type Account struct {
	ID        int
	Email     string
	Password  string
	RoleID    int
	RoleName  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
