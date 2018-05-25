package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordManager struct {

}


func (p PasswordManager)HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p PasswordManager) CompareHash(storedPass string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPass), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return false
	}
	return true
}

