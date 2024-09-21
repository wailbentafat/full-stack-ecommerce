package authentification

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"encoding/json"
	gomail "gopkg.in/mail.v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"log"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
	"io"
)

// ForgetPassword is a Gin handler that generates a random OTP
// and sends it to the user who requested a password reset.
// It also generates a JWT token for the user and sets it in
// the Gin context.
func ForgetPassword(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	log.Printf("Request body: %s", body)

	var otpRequest struct {
		Email string `json:"email"`
	}
	err = json.Unmarshal(body, &otpRequest)
	if err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if otpRequest.Email == "" {
		log.Println("No email address provided in the request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email address is required"})
		return
	}

	// Log the email that we received
	log.Printf("Received forget password request for email: %s", otpRequest.Email)
	email := otpRequest.Email

	// Generate a random OTP
	rand.Seed(time.Now().UnixNano())
	otpnum := rand.Intn(100000)
	log.Printf("Generated OTP: %d", otpnum)

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", "wailbentafat@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Forget Password")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP is %d", otpnum))

	// Create a dialer to send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, "wailbentafat@gmail.com", "bdzgoqobqdsybutc")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send OTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Sent OTP: %d", otpnum)

	// Generate a JWT token for the user
	token, err := jwt.GenerateJWT(email)
	if err != nil {
		log.Printf("Failed to generate JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Generated JWT: %s", token)

	// Return a 200 with the JWT token and the OTP
	c.JSON(http.StatusOK, gin.H{"token": token, "otp": otpnum})
}
