package routes

import (
	m"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/middleware"
	"github.com/gin-gonic/gin"
)

func SecureRoutes(router *gin.Engine) {
	router.Use(m.AuthMiddleware())
	router.GET("/secure", m.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}