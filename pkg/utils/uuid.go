package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/rs/xid"
)

func GenXID(prefix string) string {
	id := xid.New()
	return prefix + id.String()
}

func GenUUID() string {
	id := xid.New()
	return id.String()
}

func GenMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
