package authentification

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(c *gin.Context) {
	exists,ok := c.Get("email")
	if !ok{
		c.JSON(http.StatusAccepted, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"authorized": true, "email": exists})
}