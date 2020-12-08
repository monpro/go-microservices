package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	md5Instance := md5.New()
	md5Instance.Write([]byte(value))

	return hex.EncodeToString(md5Instance.Sum(nil))
}
