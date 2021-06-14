package main

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a function that converts passwords into byte, with the "golang.org/x/crypto/bcrypt" package
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckPasswordHash function allows to compare a bcrypt hashed password with ist possible plaintext equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
