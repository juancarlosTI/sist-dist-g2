package authtokenjwt

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	mappers "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/mappers_claim"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type JWTAccessTokenService struct {
	secret []byte
	ttl    time.Duration
}

func NewJWTAccessTokenService(secret []byte, ttl time.Duration) *JWTAccessTokenService {
	return &JWTAccessTokenService{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

func (j *JWTAccessTokenService) Generate(at *token_access.AccessToken) (string, error) {
	claims := mappers.ToClaims(at)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.ttl))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString(j.secret)
}

// func (j *JWTAccessTokenService) Generate(userID string, email string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"sub": userID,
// 		"email": email,
// 		"iss": "auth_service",
// 		"aud": "api",
// 		"exp": time.Now().Add(15 * time.Minute).Unix(),
// 	}

// 	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return t.SignedString(j.secret)
// }

// func (j *JWTAccessTokenService) GenerateRefreshToken(userID string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"sub": userID,
// 		"exp": time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(j.secret))
// }

func (j *JWTAccessTokenService) Validate(ctx context.Context, tokenStr string) (*token_access.AccessToken, error) {

	t, err := jwt.ParseWithClaims(tokenStr, &mappers.Claims{}, func(t *jwt.Token) (interface{}, error) {

		// valida algoritmo
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(j.secret), nil
	})

	if err != nil {
		log.Printf("JWT VALIDATION ERROR: %+v\n", err)
		log.Printf("TOKEN: %s\n", tokenStr)
		log.Printf("SECRET: %s\n", string(j.secret))
		return nil, err
	}

	log.Printf("TOKEN VALID: %+v\n", t.Valid)

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := t.Claims.(*mappers.Claims)
	if !ok {
		log.Printf("INVALID CLAIMS TYPE: %+v\n", t.Claims)
		return nil, fmt.Errorf("invalid claims type")
	}

	log.Printf("CLAIMS: %+v\n", claims)

	// AQUI É O PONTO CRÍTICO
	domain := mappers.ToDomain(*claims)

	return &domain, nil
}
