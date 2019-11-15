package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

//GetSha1 GetSha1
func GetSha1(file io.Reader) string {

	sha1 := sha1.New()

	_, err := io.Copy(sha1, file)

	if err != nil {
		return ""
	}

	return hex.EncodeToString(sha1.Sum(nil))
}
