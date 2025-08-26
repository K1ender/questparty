package utils

import (
	"crypto/rand"
)

var alphabet = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict"

func NanoID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphabet[bytes[i]&63]
	}

	return string(result), nil
}
