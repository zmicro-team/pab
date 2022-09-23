package util

import (
	"encoding/hex"
	"strings"
)

func ByteToHex(b []byte) string {
	return strings.ToUpper(hex.EncodeToString(b))
}

func HexToByte(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}
