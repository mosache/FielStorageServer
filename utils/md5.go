package utils

import "crypto/md5"

//MD5 md5
func MD5(str string) string {
	md5 := md5.New()

	cryptoBytes := md5.Sum([]byte(str))

	return string(cryptoBytes)
}
