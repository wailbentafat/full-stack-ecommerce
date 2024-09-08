package order

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type OrderResponse struct {
	Email        string `json:"email"`
	ProductTitle string `json:"product_title"`
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
}

func GetOrders(c *gin.Context) {
	
	query := `
		SELECT 
			u.email AS email,
			p.name AS product_title,
			p.price,
			c.quantity
		FROM 
			commande c
		JOIN 
			user u ON c.user_id = u.id
		JOIN 
			product p ON c.product_id = p.id
		WHERE
			c.taille = p.taille
	`
	
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to get orders: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	var orders []OrderResponse
	for rows.Next() {
		var order OrderResponse
		if err := rows.Scan(&order.Email, &order.ProductTitle, &order.Price, &order.Quantity); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}