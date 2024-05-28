package main

import (
	"gin-project/controllers"
	"gin-project/services"

	"github.com/gin-gonic/gin"
)

var (
	VideoService    services.VideoService       = services.New()
	VideoController controllers.VideoController = controllers.New(VideoService)
)

func main() {
	server := gin.Default()

	/* Endpoints */
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	/************/

	server.Run(":8080")
}
