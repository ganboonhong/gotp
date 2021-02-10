package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
