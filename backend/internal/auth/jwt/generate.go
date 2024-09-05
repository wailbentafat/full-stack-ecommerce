package jwt

import (
	"time"
	jwt "github.com/golang-jwt/jwt/v5"
	s"github.com/wailbentafat/full-stack-ecommerce/backend/pkg/config"
)


func GenerateJWT(email string) (string, error) {
	token,err :=jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":email,
			"exp":time.Now().Add(time.Hour * 100).Unix(),
		}).SignedString([]byte(s.SecretKey))
		if err!=nil{
			return "",err
		}
		return token,nil}