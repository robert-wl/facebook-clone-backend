package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (string, error) {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hashed), nil
	}
}

func ComparePassword(hashed string, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
		return false
	}
	return true
}
