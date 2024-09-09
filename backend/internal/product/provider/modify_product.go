package product

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
)



func  UpdateProduct(c*gin.Context){
	var update struct{
		Name        *string `json:"name"`
		Price       *int    `json:"price"`
		Image       *string `json:"image"`
		Description *string `json:"description"`
		Category    *string `json:"category"`
		Quantity    *int    `json:"quantity"`
	}
	id :=c.Param("id")
	if id==""{
		log.Println("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}

	if err:=c.BindJSON(&update);err!=nil{
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

   var queryset[]string
   var arguments[]interface{}
   if update.Name!=nil{
	queryset=append(queryset,"name=?")
	arguments=append(arguments,*update.Name)
   }
   if update.Price!=nil{
	queryset=append(queryset,"price=?")
	arguments=append(arguments,*update.Price)
   }
   if update.Image!=nil{
	queryset=append(queryset,"image=?")
	arguments=append(arguments,*update.Image)
   }
   if update.Description!=nil{
	queryset=append(queryset,"description=?")
	arguments=append(arguments,*update.Description)
   }
   if update.Category!=nil{
	queryset=append(queryset,"category=?")
	arguments=append(arguments,*update.Category)
   }
   if update.Quantity!=nil{
	queryset=append(queryset,"quantity=?")
	arguments=append(arguments,*update.Quantity)
   }
   if len(queryset)==0{
	log.Println("No fields to update")
	c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
	return
   }
   query:=`UPDATE product SET `+strings.Join(queryset,", ")+` WHERE id = ?`
   arguments=append(arguments,id)
   _,err:=db.Exec(query,arguments...)
   if err!=nil{
	log.Printf("Failed to update product: %v\n", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
   }

   c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}