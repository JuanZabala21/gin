package controllers

import (
	"gin-project/entities"
	"gin-project/services"
	"gin-project/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controllers struct {
	services services.VideoService
}

var validate *validator.Validate

func New(s services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controllers{
		services: s,
	}
}

func (c *controllers) FindAll() []entities.Video {
	return c.services.FindAll()
}

func (c *controllers) Save(ctx *gin.Context) error {
	var video entities.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}
	if err := validate.Struct(video); err != nil {
		return err
	}
	c.services.Save(video)
	return nil
}

func (c *controllers) ShowAll(ctx *gin.Context) {
	videos := c.services.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controllers) Update(ctx *gin.Context) error {
	var video entities.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id
	if err := validate.Struct(video); err != nil {
		return err
	}
	c.services.Update(video)
	return nil
}

func (c *controllers) Delete(ctx *gin.Context) error {
	var video entities.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}

	video.ID = id
	c.services.Delete(video)
	return nil
}
