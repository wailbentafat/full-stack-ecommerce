package product

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateProduct(c *gin.Context) {
	fmt.Println("Received request to create new product")

	// Parse the form with a 10 MB limit
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Failed to parse form: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Extract form values
	name := c.Request.FormValue("name")
	priceStr := c.Request.FormValue("price")
	description := c.Request.FormValue("description")
	category := c.Request.FormValue("category")
	quantityStr := c.Request.FormValue("quantity")

	// Extract file
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		log.Printf("Failed to get file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()

	// Convert price and quantity to integers
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		log.Printf("Invalid price value: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		log.Printf("Invalid quantity value: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	// Check for missing fields
	if name == "" || description == "" || category == "" || quantityStr == "" || file == nil {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	_, currentFilePath, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(currentFilePath)
	uploadDir := filepath.Join(baseDir, "uploads")

	
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Printf("Failed to create upload directory: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}
	file, fileHeader, err := c.Request.FormFile("image")
    if err != nil {
    log.Printf("Failed to get file: %v\n", err)
    c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
    return
}
    defer file.Close()



	fileName := uuid.New().String() + filepath.Ext(fileHeader.Filename)
	filePath := filepath.Join(uploadDir, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Failed to copy file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Successfully created and copied file")

	query := `INSERT INTO product (name, price, image, description, category, quantity) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, name, price, filePath, description, category, quantity)
	if err != nil {
		log.Printf("Failed to insert product: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Successfully inserted product")

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
