package auth

import (
	"fmt"
	"social-network/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(config.Config("JWT_SECRET_KEY"))

func GenerateJWT(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()
	claims["authorized"] = true
	claims["userId"] = userId

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("failed with JWT signing method")
		}
		return sampleSecretKey, nil
	})
	return err
}
