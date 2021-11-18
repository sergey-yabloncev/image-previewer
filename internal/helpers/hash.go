package helpers

import (
	"crypto/md5"
	"fmt"
)

func Md5String(in string) string {
	if in == "" {
		return ""
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}
