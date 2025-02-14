package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes - Register all routes calling the respective route registration functions
func RegisterRoutes(router *gin.Engine) {
	RegisterUserRoutes(router)
}
