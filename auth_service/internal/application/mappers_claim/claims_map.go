package mappers_claim

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string            `json:"sub"`
	Roles  string            `json:"roles"`
	Autor  map[string]string `json:"autor"`
	Origem map[string]string `json:"origem"`

	jwt.RegisteredClaims
}
