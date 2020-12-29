package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/gin-poc/entity"
	"gitlab.com/pragmaticreviews/gin-poc/service"
)

// controller interface: 사용할 함수를 선언
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

// controller struct : service와 연결 역할
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

// VideoController interface의 함수 선언
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// VideoController interface의 함수 선언
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video

}
