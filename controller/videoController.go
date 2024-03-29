package controller

import (
	"net/http"

	"github.com/amro-alasri/golangBasics/entity"
	"github.com/amro-alasri/golangBasics/service"
	"github.com/amro-alasri/golangBasics/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(server service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: server,
	}
}

func (s *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	s.service.Save(video)
	return nil
}

func (s *controller) FindAll() []entity.Video {
	return s.service.FindAll()
}

func (s *controller) ShowAll(ctx *gin.Context) {
	videos := s.service.FindAll()

	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)

}
