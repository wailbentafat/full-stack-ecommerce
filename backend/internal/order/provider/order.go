package order

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	
	"database/sql"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateOrder(c *gin.Context) {
	var order struct{
		user_id int `json:"user_id"`
		ProductID int `json:"product_id"`
		Quantity int `json:"quantity"`
		taille string `json:"taille"`
	}
	if err := c.BindJSON(&order); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `INSERT INTO order (user_id, product_id, quantity, taille) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, order.user_id, order.ProductID, order.Quantity, order.taille)
	if err != nil {
		log.Printf("Failed to insert order: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	query=`UPDATE product SET quantity=quantity-? WHERE taille=?`
	_, err = db.Exec(query, order.Quantity, order.taille)
	if err != nil {
		log.Printf("Failed to update product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	query=`UPDATE product SET quantity=quantity-? WHERE id=?`
	_, err = db.Exec(query, order.Quantity, order.ProductID)
	if err != nil {
		log.Printf("Failed to update product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
