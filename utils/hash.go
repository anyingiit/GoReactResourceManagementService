package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func BytesEncodingToHashSha256HexString(data []byte) string {
	hash := sha256.Sum256(data)
	hexString := hex.EncodeToString(hash[:])

	return hexString
}
