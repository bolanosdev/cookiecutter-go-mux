package utils

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts/errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("utils.HashPassword", err)
	}

	return string(hashed_password), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("utils.CheckPassword", err)
	}
	return nil
}
