package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// StringMD5V 加密
func StringMD5V(value string, b ...byte) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(b))
}

// BytesMD5V 加密
func BytesMD5V(value []byte, b ...byte) string {
	m := md5.New()
	m.Write(value)
	return hex.EncodeToString(m.Sum(b))
}
