package oidc

import (
	"context"
	"errors"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

type NormalizedClaims struct {
	Subject   string
	Email     string
	Name      string
	Issuer    string
	Audience  string
	RawClaims map[string]any
}

type OIDCProvider interface {
	Name() string
	ExchangeCode(ctx context.Context, code string, redirectURI string) (*TokenResponse, error)
	VerifyIDToken(ctx context.Context, idToken string) (*NormalizedClaims, error)
}

var (
	ErrProviderNotFound = errors.New("oidc: provider not found")
	ErrInvalidIDToken   = errors.New("oidc: invalid id token")
	ErrMissingCode      = errors.New("oidc: missing authorization code")
)
