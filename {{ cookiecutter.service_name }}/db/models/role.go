package models

type Role struct {
	ID          int
	Name        string
	Permissions []Permission
}
