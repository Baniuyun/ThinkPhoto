package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func EncPassword(password, uuid string) string {
	return Md5Sum([]byte(strings.TrimSpace(password + uuid)))
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}
