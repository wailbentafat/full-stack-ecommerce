package adminpackage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	 m"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
)
var admin bool

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
       
		
		tokenStr := parts[1]
		claims, err := m.Parsejwt(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
        if claims=="www.hamidbentafat@gmail.com"{
			admin=true
		}else{
			admin=false
		}
		
		if !admin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not an admin"})
			c.Abort()
			return
		}
		c.Set("email", claims)
		c.Set("isadmin", admin)
		c.Next()
	}
}
