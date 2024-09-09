package routes

import (
	m"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/middleware"
	"github.com/gin-gonic/gin"
"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/authentification")

func SecureRoutes(router *gin.Engine) {
	router.Use(m.AuthMiddleware())
	router.GET("/authorized",authentification.Check)
}
