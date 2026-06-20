package oidc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleProviderConfig struct {
	ClientID      string
	ClientSecret  string
	Issuer        string
	TokenEndpoint string
	JWKSURI       string
	HTTPTimeout   time.Duration
	JWKSCacheTTL  time.Duration
}

type GoogleProvider struct {
	config GoogleProviderConfig
	jwks   *JWKSCache
	client *http.Client
}

func NewGoogleProvider(cfg GoogleProviderConfig) (*GoogleProvider, error) {
	if cfg.ClientID == "" {
		return nil, errors.New("google oidc provider requires client id")
	}
	if cfg.ClientSecret == "" {
		return nil, errors.New("google oidc provider requires client secret")
	}
	if cfg.Issuer == "" {
		cfg.Issuer = "https://accounts.google.com"
	}
	if cfg.TokenEndpoint == "" {
		cfg.TokenEndpoint = "https://oauth2.googleapis.com/token"
	}
	if cfg.JWKSURI == "" {
		cfg.JWKSURI = "https://www.googleapis.com/oauth2/v3/certs"
	}
	if cfg.HTTPTimeout == 0 {
		cfg.HTTPTimeout = 10 * time.Second
	}
	if cfg.JWKSCacheTTL == 0 {
		cfg.JWKSCacheTTL = 10 * time.Minute
	}

	return &GoogleProvider{
		config: cfg,
		jwks:   NewJWKSCache(cfg.JWKSURI, cfg.JWKSCacheTTL, &http.Client{Timeout: cfg.HTTPTimeout}),
		client: &http.Client{Timeout: cfg.HTTPTimeout},
	}, nil
}

func (p *GoogleProvider) Name() string {
	return "google"
}

func (p *GoogleProvider) ExchangeCode(ctx context.Context, code string, redirectURI string) (*TokenResponse, error) {
	if code == "" {
		return nil, ErrMissingCode
	}
	if redirectURI == "" {
		return nil, errors.New("oidc: missing redirect URI")
	}

	values := url.Values{}
	values.Set("code", code)
	values.Set("client_id", p.config.ClientID)
	values.Set("client_secret", p.config.ClientSecret)
	values.Set("redirect_uri", redirectURI)
	values.Set("grant_type", "authorization_code")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.config.TokenEndpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("oidc: token exchange failed (%d): %s", resp.StatusCode, string(body))
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	if tokenResponse.IDToken == "" {
		return nil, errors.New("oidc: id_token missing in token response")
	}

	return &tokenResponse, nil
}

func (p *GoogleProvider) VerifyIDToken(ctx context.Context, idToken string) (*NormalizedClaims, error) {
	if idToken == "" {
		return nil, ErrInvalidIDToken
	}

	claims := &IDTokenClaims{}
	token, err := jwt.ParseWithClaims(
		idToken,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("oidc: unsupported signing method %v", t.Header["alg"])
			}

			kid, ok := t.Header["kid"].(string)
			if !ok || kid == "" {
				return nil, errors.New("oidc: missing token kid header")
			}

			return p.jwks.GetKey(ctx, kid)
		},
		jwt.WithIssuer(p.config.Issuer),
		jwt.WithAudience(p.config.ClientID),
		jwt.WithLeeway(5*time.Minute),
	)

	if err != nil {
		return nil, fmt.Errorf("oidc: id token validation failed: %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidIDToken
	}

	if claims.Subject == "" || claims.Email == "" {
		return nil, errors.New("oidc: id token missing required claims")
	}

	return claims.ToNormalizedClaims(), nil
}
