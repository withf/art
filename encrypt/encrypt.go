package encrypt

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

func Encrypt(b []byte) string {
	mb := md5.Sum(b)
	sb := sha512.Sum384(mb[:])
	return hex.EncodeToString(sb[:])
}
