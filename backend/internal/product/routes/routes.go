package product_routes


import (
	"github.com/gin-gonic/gin"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/middleware"
	"database/sql"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/core/cach"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/provider"
	
)

func Routes(router *gin.Engine,db *sql.DB,filecache *cach.Filecach) {
	product.SetDB(db)
	product.Setcach(filecache)
	router.GET("/products/getall",product.GetAllProduct)
	router.GET("/product/:id", product.GetProduct)
	router.POST("/product", a.AdminMiddleware(),product.CreateProduct)
	router.PUT("/product/:id", a.AdminMiddleware(),product.UpdateProduct)
	router.DELETE("/product/:id",a.AdminMiddleware(), product.DeleteProduct)
}