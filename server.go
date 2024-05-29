package main

import (
	"gin-project/controllers"
	"gin-project/middlewares"
	"gin-project/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	/* Endpoints */
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		if err := VideoController.Save(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video is Valid!"})
		}
	})
	/************/

	server.Run(":8080")
}
