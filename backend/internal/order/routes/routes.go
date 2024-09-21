package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/full-stack-ecommerce/backend/internal/order/provider"
	a"github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/middleware"
  
)
func Routes(router *gin.Engine) {
router.POST("/order", a.AuthMiddleware(),order.PlaceOrder)	
}