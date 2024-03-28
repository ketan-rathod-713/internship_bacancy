package controller

import (
	"ginbasics/entity"
	"ginbasics/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

var validate *validator.Validate

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	// we can bind query, uri, yaml and headers and so on...
	// BindJson is a shortcut for mustBindWith

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func New(service service.VideoService) VideoController {
	validate = validator.New()

	// add validation tag with given key.
	validate.RegisterValidation("is-cool", nil)
	return &controller{ // why giving error, due to because of pointer receiver may be.
		service: service,
	}
}
