package product_routes


import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/provider"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/middleware"
	"database/sql"
)

func Routes(router *gin.Engine,db *sql.DB) {
	product.SetDB(db)
	router.GET("/products/getall",product.GetAllProduct)
	router.GET("/product/:id", product.GetProduct)
	router.POST("/product", a.AdminMiddleware(),product.CreateProduct)
	router.PUT("/product/:id", a.AdminMiddleware(),product.UpdateProduct)
	router.DELETE("/product/:id",a.AdminMiddleware(), product.DeleteProduct)
}