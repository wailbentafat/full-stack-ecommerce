package jwt

import (
	
	jwt "github.com/golang-jwt/jwt/v5"
	"backend/pkg/config"
)

func Parsejwt(token string) (string, error) {
	claims := jwt.MapClaims{}
	_,err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.config("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	return claims["email"].(string), nil
}