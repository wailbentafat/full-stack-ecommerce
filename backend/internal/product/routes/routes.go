package product_routes


import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/provider"
)

func Routes(router *gin.Engine) {
	router.GET("/products", product.GetAllProduct)
	router.GET("/product/:id", product.GetProduct)
	router.POST("/product", product.CreateProduct)
	router.PUT("/product/:id", product.UpdateProduct)
	router.DELETE("/product/:id", product.DeleteProduct)
}