package authentification

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	j"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := user.Email
	password := user.Password

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO user (email, password) VALUES (?, ?)`
	_, err = db.Exec(query, email, string(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := j.GenerateJWT(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}