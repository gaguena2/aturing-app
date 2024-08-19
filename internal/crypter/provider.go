package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type ProviderCrypter struct {
	key string
}

func NewCrypter(key string) *ProviderCrypter {
	return &ProviderCrypter{key}
}

func (c *ProviderCrypter) Encrypt(plaintext []byte) []byte {
	aes, err := aes.NewCipher([]byte(c.key))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return ciphertext
}

func (c *ProviderCrypter) Decrypt(ciphertext string) []byte {
	aes, err := aes.NewCipher([]byte(c.key))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}
