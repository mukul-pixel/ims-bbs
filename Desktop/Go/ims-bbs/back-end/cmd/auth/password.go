package auth

import "golang.org/x/crypto/bcrypt"

func HashThePassword(password string) (string, error) {
	hashed,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}

	return string(hashed),nil
}