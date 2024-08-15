package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type defaultClaims struct {
	jwt.RegisteredClaims
	Roles []uint64
}

func NewJWT(sub, id string, days int, secret, issuer string, roles []uint64) (string, error) {
	now := time.Now()
	claim := defaultClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   sub,
			ID:        id,
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, days)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Roles: roles,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
}

func ParseJWT(token, secret string) (*defaultClaims, error) {
	t, err := jwt.ParseWithClaims(token, &defaultClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*defaultClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
