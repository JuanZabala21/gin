package controllers

import (
	"gin-project/entities"
	"gin-project/services"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) entities.Video
}

type controllers struct {
	services services.VideoService
}

func New(s services.VideoService) VideoController {
	return &controllers{
		services: s,
	}
}

func (c *controllers) FindAll() []entities.Video {
	return c.services.FindAll()
}

func (c *controllers) Save(ctx *gin.Context) entities.Video {
	var video entities.Video
	ctx.BindJSON(&video)
	c.services.Save(video)
	return video
}
