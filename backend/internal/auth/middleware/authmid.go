package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	 m"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
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

		
		c.Set("email", claims)

		
		c.Next()
	}
}
