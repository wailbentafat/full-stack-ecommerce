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
type order struct{
	productid int  `json:"productid"`
	taille string `json:"taille"`
	quantity int   `json:"quantity"`
}
func PlaceOrder(c *gin.Context) {
type OrderRequest struct {
	Email string `json:"email"`
	orders[]order `json:"orders"`}
	var Orderrequest OrderRequest
if err := c.BindJSON(&Orderrequest); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
b,err:=db.Begin()
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})	
	return

}
defer func() {
	if err!=nil{
		b.Rollback()
	}else{
		b.Commit()
	}
}()
var user_id int
var total int
var price int
var exists bool

total=0
err=b.QueryRow("SELECT id FROM user WHERE email = ?",Orderrequest.Email).Scan(&user_id)
if err != nil {
	b.Rollback()
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	return 
}
for _,order:=range Orderrequest.orders{
  err=b.QueryRow("SELECT EXISTS(SELECT 1 FROM product WHERE id = ?)",order.productid).Scan(&exists)
  if err != nil || !exists {
	b.Rollback()
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return 
  }
  err=b.QueryRow("SELECT price FROM product WHERE id = ?",order.productid).Scan(&price)
  total+=price*order.quantity
  if err != nil {
	b.Rollback()
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	  return}
  var quantity int
  err:=b.QueryRow("SELECT quantity FROM taille WHERE taille = ?",order.taille).Scan(&quantity)
  if quantity<order.quantity{
	b.Rollback()
	c.JSON(http.StatusInternalServerError, gin.H{"error": "end of stock"})
	return 
  }
  _,err=b.Exec("INSERT INTO commande (user_id,product_id,quantity,taille,price,date) VALUES (?,?,?,?,?,CURRENT_TIMESTAMP)",user_id,order.productid,order.quantity,order.taille,price)
  if err != nil {
	b.Rollback()
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
  }
  _,err=b.Exec("UPDATE taille SET quantity = quantity - ? WHERE taille = ?",order.quantity,order.taille)

  if err != nil {
	b.Rollback()
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
  }
  _,err=b.Exec("UPDATE product SET quantity = quantity - ? WHERE id = ?",order.quantity,order.productid)

  if err != nil {
	b.Rollback()
	  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
  }
  
  c.JSON(http.StatusOK, gin.H{"message": "order placed successfully","total":total})

}
}