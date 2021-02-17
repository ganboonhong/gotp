package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return hash
}
