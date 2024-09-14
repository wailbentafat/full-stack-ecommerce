package authentification
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/jwt"
)


var (
	oauth2Config = &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
	oauth2State = "kegfwegegfkeekfgekflwfL"
)

func loginoauth(c*gin.Context) {
	url := oauth2Config.AuthCodeURL(oauth2State)
	c.Redirect(http.StatusTemporaryRedirect, url)

}

func Callback(c*gin.Context) {
	
	state := c.Query("state")
	if state != oauth2State {
		fmt.Println("Invalid state parameter")
		return
	}
	code := c.Query("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Failed to exchange token:", err)
		return
	}
	
	client:= oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		fmt.Println("Failed to get user info:", err)
		return
	}
	defer resp.Body.Close()
	
	var profile map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		fmt.Println("Failed to decode response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}
	
	fmt.Println("Profile:", profile)
	query:=`INSERT INTO user (email, password) VALUES (?, ?)`
	_, err = db.Exec(query, profile["email"], 000000)
	if err != nil {
		fmt.Println("Failed to insert user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}
	acctoken, err := jwt.GenerateJWT(profile["email"].(string))
	if err != nil {
		fmt.Println("Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": acctoken})
}