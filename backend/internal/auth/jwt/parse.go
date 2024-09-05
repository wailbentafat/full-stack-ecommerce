package jwt

import (
	
	jwt "github.com/golang-jwt/jwt/v5"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/pkg/config"
)

func Parsejwt(token string) (string, error) {
	claims := jwt.MapClaims{}
	_,err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	return claims["email"].(string), nil
}