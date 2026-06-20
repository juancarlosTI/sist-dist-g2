package token_access

import "time"

type RefreshToken struct {
	Value     string
	UserID    string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}

func (rt *RefreshToken) IsExpired(now time.Time) bool {
	return now.After(rt.ExpiresAt)
}

func (rt *RefreshToken) IsRevoked() bool {
	return rt.Revoked
}

func (rt *RefreshToken) Revoke() {
	rt.Revoked = true
}
