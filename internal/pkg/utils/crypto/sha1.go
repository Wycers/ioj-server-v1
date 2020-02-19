package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(src string) string {
	h := sha1.New()
	h.Write([]byte(src))

	return hex.EncodeToString(h.Sum([]byte("")))
}
