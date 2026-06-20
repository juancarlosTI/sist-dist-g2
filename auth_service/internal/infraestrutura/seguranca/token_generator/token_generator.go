package token_generator

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type SecureTokenGenerator struct {
	size int // tamanho em bytes
}

func NewSecureTokenGenerator(size int) token_access.TokenGenerator {
	return &SecureTokenGenerator{
		size: size,
	}
}

func (g *SecureTokenGenerator) Generate() (string, error) {
	b := make([]byte, g.size)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
