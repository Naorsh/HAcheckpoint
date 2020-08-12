package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

const key string = "301119892014882281234567"

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt - encryps app key to store securly in DB
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Failed to Encrypt")
		return "", err
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext), nil
}

// Decrypt - decrypts app key for presentation
func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Failed to Decrypt")
		return "", err
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
