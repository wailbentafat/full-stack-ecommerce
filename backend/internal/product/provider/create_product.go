package product

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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



	err := c.Request.ParseMultipartForm(10 << 20) 
    if err != nil {
        log.Printf("Failed to parse form: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
        return
    }



    file, _, err := c.Request.FormFile("image")
    if err != nil {
        log.Printf("Failed to get file: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
        return
    }
    defer file.Close()



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
	fileName := uuid.New().String() + filepath.Ext(c.Request.FormValue("image"))
    filePath := "./uploads/" + fileName
	out, err := os.Create(filePath)
	if err!=nil{
		log.Printf("Failed to create file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err!=nil{
		log.Printf("Failed to copy file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	
	query := `INSERT INTO product (name, price, image, description, category, quantity) VALUES (?, ?, ?, ?, ?, ?)`
	_ , err = db.Exec(query, product.Name, product.Price,filePath, product.Description, product.Category, product.Quantity)
	if err != nil {
		log.Printf("Failed to insert product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
