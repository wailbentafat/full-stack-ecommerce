package authentification

import (
	"database/sql"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	
	j "github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}
func Login(c *gin.Context) {
	var Loginval struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&Loginval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := Loginval.Email
	password := Loginval.Password

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	
	var storedHash string
	query := `SELECT password FROM user WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&storedHash)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := j.GenerateJWT(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
