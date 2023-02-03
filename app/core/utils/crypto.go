package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"maply/config"
	"strconv"
	"time"
)

// HashPassword Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a hex string
func HashPassword(password string) string {
	var passwordBytes = []byte(password)
	passwordBytes = append(passwordBytes, config.C.Auth.PasswordSalt...)

	var hash = sha512.New()
	hash.Write(passwordBytes)

	hashedPasswordBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedPasswordBytes)
}

func HashFileName(s string) string {
	var fileBytes = []byte(s)
	var now = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	fileBytes = append(fileBytes, now...)

	var hash = sha256.New()
	hash.Write(fileBytes)

	hashedFileBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedFileBytes)
}
