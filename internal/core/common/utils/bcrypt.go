package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func BcryptGen(str string) string {
	password := []byte(str)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	return string(hash)
}

func BcryptVerify(strPass, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(strPass))
}
