package hash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type SHA256TokenHasher struct{}

func NewSHA256TokenHasher() token_access.TokenHasher {
	return &SHA256TokenHasher{}
}

func (h *SHA256TokenHasher) Hash(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}
