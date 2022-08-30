package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(text string) string {
	data := []byte(text)
	encoded := md5.Sum(data)
	return hex.EncodeToString(encoded[:])
}
