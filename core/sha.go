package core

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
