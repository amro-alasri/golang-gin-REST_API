package controller

import (
	"github.com/amro-alasri/golangBasics/entity"
	"github.com/amro-alasri/golangBasics/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(server service.VideoService) VideoController {
	return &controller{
		service: server,
	}
}

func (s *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	s.service.Save(video)
	return video
}

func (s *controller) FindAll() []entity.Video {
	return s.service.FindAll()
}
