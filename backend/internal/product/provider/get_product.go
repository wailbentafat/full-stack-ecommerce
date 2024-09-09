package product


import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetProduct(c*gin.Context){
	id :=c.Param("id")
	if id == "" {
		log.Println("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}

	query := `SELECT * FROM product WHERE id = ?`
	row := db.QueryRow(query, id)
	var product Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.Description, &product.Category, &product.Quantity); err != nil {
		log.Printf("Failed to get product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println(product)
	c.JSON(http.StatusOK, product)
}