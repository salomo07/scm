package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 10
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		print("\n" + err.Error() + "\n")
		return "", err
	}

	// Combine the salt and hashed password
	hashedPassword := string(salt)

	return hashedPassword, nil
}
