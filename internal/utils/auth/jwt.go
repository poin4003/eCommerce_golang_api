package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/poin4003/eCommerce_golang_api/global"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	// 1. Set time expiration
	timeEx := global.Config.JWT.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}
