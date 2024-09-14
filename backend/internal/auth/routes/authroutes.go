package routes

import (
	"database/sql"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/authentification"
	"github.com/gin-gonic/gin"

)

func AuthRoutes(router *gin.Engine, db *sql.DB) {
	authentification.SetDB(db)
	router.POST("/register", authentification.Register)
	router.POST("/login", authentification.Login)
	router.POST("/forget-password", authentification.ForgetPassword)
	router.GET("/auth/callback", authentification.Callback)
	
}

