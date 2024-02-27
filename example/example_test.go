package example_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stuckinforloop/jwt-codegen/example"
)

func TestToken(t *testing.T) {
	s := jwt.StandardClaims{ExpiresAt: time.Now().Add(5 * time.Minute).UnixNano()}
	c := example.CustomClaims{Username: "stuckinforloop", StandardClaims: s}
	tokenString, err := example.GenerateToken(c)
	if err != nil {
		t.Error(err)
	}
	_, err = example.ValidateToken(tokenString)
	if err != nil {
		t.Error(err)
	}
}
