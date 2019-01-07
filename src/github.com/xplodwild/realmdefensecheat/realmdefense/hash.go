package realmdefense

import (
	"crypto/md5"
	"fmt"
)

var kSecret = []byte{0xE7, 0xA7, 0x98, 0xE5, 0xAF, 0x86}

// ComputeHash computes the HTTP Header "Hash". Hello, Crypto.ComputeHash ;-)
func ComputeHash(body []byte) string {
	h := md5.New()
	h.Write(body)
	h.Write(kSecret)
	return fmt.Sprintf("%x", h.Sum(nil))
}
