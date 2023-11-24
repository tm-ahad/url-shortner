package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func RandString() string {
	b := make([]byte, 8);
	_, e := rand.Read(b)

	HandleErr(e)

	return hex.EncodeToString(b)
}
