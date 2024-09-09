package product
import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Failed to delete file: %v\n", err)
		return err
	}
	return nil
}

func DeleteProduct(c *gin.Context) {
	

	id := c.Param("id")
    
	if id == "" {
		log.Println("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}
    var image string
	err := db.QueryRow("SELECT image FROM product WHERE id = ?", id).Scan(&image)	
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			log.Printf("Failed to fetch product: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		}
		
	
    
	err = DeleteFile(image)
	if err != nil {
		log.Printf("Failed to delete file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}
	_, err = db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		log.Printf("Failed to delete product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}