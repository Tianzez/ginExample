package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(raw string) string {
	hashed := md5.Sum([]byte(raw))
	return fmt.Sprintf("%x", hashed)
}
