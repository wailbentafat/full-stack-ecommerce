package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/order/provider"
	m"github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/middleware"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/middleware"
  
)
func Routes(router *gin.Engine) {
	router.GET("/orders",m.AdminMiddleware(),order.GetOrders)
	router.POST("/place-order",a.AuthMiddleware(),order.CreateOrder)
	
}