package routes

import (
	"github.com/gin-gonic/gin"
)

func RunRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/api/user")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"route_name": router.BasePath(),
			"message":    "",
		})
	})
}
