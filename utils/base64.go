package utils

import (
	"encoding/base64"
)

func BytesToBase64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64ToBytes(message string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(message)
}
