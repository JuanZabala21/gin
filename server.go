package main

import (
	"gin-project/controllers"
	"gin-project/middlewares"
	"gin-project/services"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	VideoService    services.VideoService       = services.New()
	VideoController controllers.VideoController = controllers.New(VideoService)
)

func setupLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutPut()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

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
