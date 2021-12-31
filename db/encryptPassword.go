package db

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword is a routine that allows me to encrypt the password */
func EncryptPassword(pass string) (string, error) {
	costo := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
