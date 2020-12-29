package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/gin-poc/controller"
	"gitlab.com/pragmaticreviews/gin-poc/service"
)

// 사용할 변수를 묶어서 밖에 선언
var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	// video 정보 get
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	// video 정보 post
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":7070")
}
