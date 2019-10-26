package util

import (
	"encoding/base64"
	"crypto/rand"
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
