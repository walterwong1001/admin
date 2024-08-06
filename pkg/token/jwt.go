package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewJWT(sub, id string, days int, secret, issuer string) (string, error) {
	now := time.Now()
	claim := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   sub,
		ID:        id,
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, days)),
		IssuedAt:  jwt.NewNumericDate(now),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
}
