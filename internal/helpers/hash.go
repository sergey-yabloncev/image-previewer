package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func Md5String(in string) string {
	if in == "" {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(in))

	return hex.EncodeToString(h.Sum(nil))
}
