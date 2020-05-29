package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"pcc.com/golangTest/golang-gin-poc/entity"
	"pcc.com/golangTest/golang-gin-poc/service"
	"pcc.com/golangTest/golang-gin-poc/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	vedioService service.IVideoService
}

var validate *validator.Validate

func New(service service.IVideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		vedioService: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.vedioService.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.vedioService.Save(video)
	return nil
}
