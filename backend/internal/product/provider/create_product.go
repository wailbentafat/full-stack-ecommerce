package product

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateProduct(c *gin.Context) {
	var product struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Quantity    int    `json:"quantity"`
	}

	if err := c.BindJSON(&product); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if product.Name == "" || product.Price == 0 || product.Image == "" || product.Description == "" || product.Category == "" || product.Quantity == 0 {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	
	query := `INSERT INTO product (name, price, image, description, category, quantity) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, product.Name, product.Price, product.Image, product.Description, product.Category, product.Quantity)
	if err != nil {
		log.Printf("Failed to insert product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
