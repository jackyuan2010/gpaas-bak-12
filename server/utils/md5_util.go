package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}