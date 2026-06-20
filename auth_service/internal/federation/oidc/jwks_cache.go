package oidc

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

type jwkKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use,omitempty"`
	Alg string `json:"alg,omitempty"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type jwksResponse struct {
	Keys []jwkKey `json:"keys"`
}

type JWKSCache struct {
	url       string
	ttl       time.Duration
	createdAt time.Time
	keys      map[string]*rsa.PublicKey
	client    *http.Client
	mu        sync.RWMutex
}

func NewJWKSCache(url string, ttl time.Duration, client *http.Client) *JWKSCache {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	return &JWKSCache{
		url:    url,
		ttl:    ttl,
		keys:   map[string]*rsa.PublicKey{},
		client: client,
	}
}

func (c *JWKSCache) GetKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	if kid == "" {
		return nil, errors.New("oidc: missing key id")
	}

	c.mu.RLock()
	if time.Since(c.createdAt) < c.ttl {
		if key, ok := c.keys[kid]; ok {
			c.mu.RUnlock()
			return key, nil
		}
	}
	c.mu.RUnlock()

	if err := c.refresh(ctx); err != nil {
		return nil, err
	}

	c.mu.RLock()
	key, ok := c.keys[kid]
	c.mu.RUnlock()
	if !ok {
		return nil, errors.New("oidc: jwks key not found")
	}

	return key, nil
}

func (c *JWKSCache) refresh(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("oidc: failed to fetch jwks")
	}

	var payload jwksResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return err
	}

	keys := map[string]*rsa.PublicKey{}
	for _, jwk := range payload.Keys {
		if jwk.Kty != "RSA" {
			continue
		}

		publicKey, err := parseRSAPublicKey(jwk)
		if err != nil {
			return err
		}

		keys[jwk.Kid] = publicKey
	}

	c.mu.Lock()
	c.keys = keys
	c.createdAt = time.Now()
	c.mu.Unlock()

	return nil
}

func parseRSAPublicKey(jwk jwkKey) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(jwk.N))
	if err != nil {
		return nil, err
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(jwk.E))
	if err != nil {
		return nil, err
	}

	if len(eBytes) == 0 {
		return nil, errors.New("oidc: invalid exponent value")
	}

	publicExponent := 0
	for _, b := range eBytes {
		publicExponent = publicExponent<<8 + int(b)
	}

	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: publicExponent,
	}, nil
}
