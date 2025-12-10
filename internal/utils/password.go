package utils

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

const BcryptCost = 12

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func ParseUserIDFromSubject(subject string) (uint, error) {
	userID, err := strconv.ParseUint(subject, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID in subject: %w", err)
	}
	return uint(userID), nil
}
