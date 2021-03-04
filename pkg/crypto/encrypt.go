package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// createHash generates fixed length (32 characters) of data
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	// Sum will return 128-bit(16 bytes) MD5 hashes
	s := hasher.Sum(nil)

	// hex.EncodeToString takes 16 bytes of data and returns 32 characters (hexadecimal code)
	return hex.EncodeToString(s)
}

// Encrypt encrypts plaintext and  returns encrypted hexadecimal string
func Encrypt(plaintext string, key string) string {
	h := createHash(key)
	block, err := aes.NewCipher([]byte(h))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	// read len(nonce) bytes from rander buffer (rand.Reader) to nonce
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext)
}

// Decrypt decrypts given encrypted hexadecimal string and returns plaintext
func Decrypt(encrypted string, key string) string {
	bytes, err := hex.DecodeString(encrypted)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher([]byte(createHash(key)))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	// we prefixed our encrypted data with the nonce. This means that we need to separate the nonce and the encrypted data.
	nonce, ciphertext := bytes[:nonceSize], bytes[nonceSize:]
	text, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return string(text)
}
