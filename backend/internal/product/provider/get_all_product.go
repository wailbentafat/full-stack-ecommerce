package product

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Product struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Price       int    `json:"price"`
    Image       string `json:"image"`
    Description string `json:"description"`
    Category    string `json:"category"`
    Quantity    int    `json:"quantity"`
}
func GetAllProduct(c *gin.Context) {
	query:=`SELECT * From product`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to get products: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	var products []Product
	  for rows.Next(){
		var p Product
		if err:=rows.Scan(&p.ID,&p.Name,&p.Price,&p.Image,&p.Description,&p.Category,&p.Quantity);err!=nil{
			log.Printf("Failed to get products: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
        log.Println(p)
		products=append(products,p)
	  }

	c.JSON(http.StatusOK, products)
}
