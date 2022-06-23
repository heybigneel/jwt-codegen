package example

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secretKey")

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(c CustomClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = c.Username
	claims["expiresAt"] = c.StandardClaims.ExpiresAt

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error parsing token")
		}
		return secretKey, nil
	})
	if err != nil {
		err = errors.New("token has expired")
		return jwt.MapClaims{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("error getting claims")
	}
	return claims, nil
}
