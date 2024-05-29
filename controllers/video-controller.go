package controllers

import (
	"gin-project/entities"
	"gin-project/services"
	"gin-project/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) error
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
