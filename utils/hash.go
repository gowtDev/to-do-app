package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword returns the SHA-256 hash of the password as a hex string
func HashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

// CheckPasswordHash compares a SHA-256 hashed password with a plain password
func CheckPasswordHash(password, hashedPassword string) bool {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]) == hashedPassword
}
