package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verifier(token string) (string, error) {
	//检查jwttoken格式
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}
	claims, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token not valid")
	}

	//检查token 过期等情况
	err = claims.Valid()
	if err != nil {
		return "", fmt.Errorf("token not valid :%v", err)
	}
	return claims.Subject, nil
}
