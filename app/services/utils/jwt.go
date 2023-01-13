package utils

import (
	"github.com/golang-jwt/jwt"
	"maply/config"
	"time"
)

func GenerateJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userId,
	})
	return token.SignedString([]byte(config.C.Auth.SigningKey))
}

func ParseToken(authToken string) (string, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(config.C.Auth.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userId := claims["sub"].(string)
	return userId, nil
}
