package router

import (
	"github.com/gin-gonic/gin"
)

func GenerateRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/api/user")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"route_name": router.BasePath(),
			"message":    "",
		})
	})
}
