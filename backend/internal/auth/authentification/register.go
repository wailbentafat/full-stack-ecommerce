package authentification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	j"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(c *gin.Context) {
	var user struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&user); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := user.Email
	password := user.Password

	if email == "" || password == "" {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to generate password hash: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO user (email, password) VALUES (?, ?)`
	_, err = db.Exec(query, email, string(hash))
	if err != nil {
		log.Printf("Failed to insert user into database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := j.GenerateJWT(email)
	if err != nil {
		log.Printf("Failed to generate JWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    log.Println("User registered successfully")
	log.Println(gin.H{"token": token})
	c.JSON(http.StatusOK, gin.H{"token": token})
}
