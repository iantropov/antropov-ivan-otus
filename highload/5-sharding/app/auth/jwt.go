package auth

import (
	"fmt"
	"social-network-5/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(config.Config("JWT_SECRET_KEY"))

func GenerateJWT(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(365 * 24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["userId"] = userId

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("failed with JWT signing method")
		}
		return sampleSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT Token")
	}
}
