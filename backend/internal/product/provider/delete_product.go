package product
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DeleteProduct(c *gin.Context) {
	

	id := c.Param("id")
	if id == "" {
		log.Println("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}

	query := `DELETE FROM product WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}