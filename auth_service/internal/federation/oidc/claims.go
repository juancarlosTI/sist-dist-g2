package oidc

import (
	"github.com/golang-jwt/jwt/v5"
)

type IDTokenClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func (c *IDTokenClaims) ToNormalizedClaims() *NormalizedClaims {
	audience := ""
	if len(c.Audience) > 0 {
		audience = c.Audience[0]
	}

	return &NormalizedClaims{
		Subject:  c.Subject,
		Email:    c.Email,
		Name:     c.Name,
		Issuer:   c.Issuer,
		Audience: audience,
		RawClaims: map[string]any{
			"email": c.Email,
			"name":  c.Name,
			"sub":   c.Subject,
			"iss":   c.Issuer,
			"aud":   c.Audience,
		},
	}
}
