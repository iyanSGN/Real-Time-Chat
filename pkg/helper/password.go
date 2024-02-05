package helpers

import (
	"realtime/app/auth"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(data *auth.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.Password = string(hashedPassword)
	return nil
}