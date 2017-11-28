package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sign(input string) string {
	h := sha256.New()
	_, err := h.Write([]byte(input))
	CheckErr(err)

	result := hex.EncodeToString(h.Sum(nil))
	return result
}
