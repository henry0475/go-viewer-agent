package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

// ToSha1 will return a string that was hashed by SHA1
func ToSha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
