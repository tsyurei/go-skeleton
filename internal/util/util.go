package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
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
