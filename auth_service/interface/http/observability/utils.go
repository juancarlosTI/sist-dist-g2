package observability

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Hash(value string) string {
	h := sha256.Sum256([]byte(value))
	return hex.EncodeToString(h[:])
}

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "***"
	}

	name := parts[0]
	if len(name) <= 2 {
		return "***@" + parts[1]
	}

	return name[:2] + "***@" + parts[1]
}
