package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/core/cach"
	"time"
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

var filecache *cach.Filecach
func Setcach(cache *cach.Filecach){
	filecache = cache
}
func GetAllProduct(c *gin.Context) {
	cachkey := "all_products"
	if data, ok := filecache.GEtkey(cachkey); ok {
		c.JSON(http.StatusOK, data)
		return
	}
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
	  productdata,err:=json.Marshal(products)
	  if err!=nil{
		  log.Printf("Failed to get products: %v\n", err)
		  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		  return
	  }
	  err=filecache.Setkey(cachkey,productdata,time.Minute*60)
	  if err!=nil{
		  log.Printf("Failed to get products: %v\n", err)
		  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		  return
	  }


	c.JSON(http.StatusOK, products)
}
