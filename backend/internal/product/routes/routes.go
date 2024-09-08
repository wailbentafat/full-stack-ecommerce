package product_routes


import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/provider"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/middleware"
)

func Routes(router *gin.Engine) {
	router.GET("/products",product.GetAllProduct)
	router.GET("/product/:id", product.GetProduct)
	router.POST("/product",a.AdminMiddleware(), product.CreateProduct)
	router.PUT("/product/:id", a.AdminMiddleware(),product.UpdateProduct)
	router.DELETE("/product/:id",a.AdminMiddleware(), product.DeleteProduct)
}