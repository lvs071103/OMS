package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"

	"github.com/cloudflare/cfssl/log"
	"golang.org/x/crypto/pbkdf2"
)

func GenerateRandomString() string {
	length := 32
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	l := len(chars)
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		log.Warningf("Error reading random bytes: %v", err)
	}
	for i := 0; i < length; i++ {
		result[i] = chars[int(result[i])%l]
	}
	return string(result)
}

// Encrypt encrypts the content with salt
func Encrypt(content string, salt string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(content), []byte(salt), 4096, 16, sha1.New))
}

func main() {
	randomString := GenerateRandomString()
	log.Infof("Generated random string: %s", randomString)
	encryptedContent := Encrypt("Jack@080716", randomString)
	log.Infof("Encrypted content: %s", encryptedContent)
}
