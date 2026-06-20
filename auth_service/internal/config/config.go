package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RabbitURL string
	HTTPPort  string

	JWTSecret []byte

	AccessTokenTTL time.Duration

	OIDCGoogleClientID     string
	OIDCGoogleClientSecret string
	OIDCGoogleIssuer       string
	OIDCGoogleTokenURL     string
	OIDCGoogleJWKSURL      string
	HTTPTimeoutSeconds     int
}

func Load() *Config {

	timeout := 10
	if v := os.Getenv("HTTP_TIMEOUT_SECONDS"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			timeout = parsed
		}
	}

	// 🔐 JWT SECRET
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 32 {
		panic("JWT_SECRET must be at least 32 characters")
	}

	// ⏱ TTL
	ttlSeconds := 900 // default 15 min
	if v := os.Getenv("ACCESS_TOKEN_TTL"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			ttlSeconds = parsed
		}
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		RabbitURL: os.Getenv("RABBIT_URL"),
		HTTPPort:  os.Getenv("HTTP_PORT"),

		JWTSecret:      []byte(secret),
		AccessTokenTTL: time.Duration(ttlSeconds) * time.Second,

		OIDCGoogleClientID:     os.Getenv("OIDC_GOOGLE_CLIENT_ID"),
		OIDCGoogleClientSecret: os.Getenv("OIDC_GOOGLE_CLIENT_SECRET"),
		OIDCGoogleIssuer:       os.Getenv("OIDC_GOOGLE_ISSUER"),
		OIDCGoogleTokenURL:     os.Getenv("OIDC_GOOGLE_TOKEN_URL"),
		OIDCGoogleJWKSURL:      os.Getenv("OIDC_GOOGLE_JWKS_URL"),
		HTTPTimeoutSeconds:     timeout,
	}
}

func (c *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
