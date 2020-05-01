package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
)

// GenerateRandomString generate a random string based on given length
func GenerateRandomString(length int) string {
	rb := make([]byte, length)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Println(err)
	}

	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}

// EncryptWitSecretKey data with app secret key
func EncryptWitSecretKey(data []byte) ([]byte, error) {
	passphrase := viper.GetString("secretKey")
	return Encrypt(data, passphrase)
}

func Encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(passphrase[:32]))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}

// DecryptWithSecretKey data with app secret key
func DecryptWithSecretKey(data []byte) ([]byte, error) {
	passphrase := viper.GetString("secretKey")
	return Decrypt(data, passphrase)
}

func Decrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(passphrase[:32]))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// DeepCopy a copier function based on same struct name https://github.com/jinzhu/copier
func DeepCopy(from, to interface{}) error {
	return copier.Copy(to, from)
}
