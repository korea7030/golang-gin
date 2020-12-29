package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"gitlab.com/pragmaticreviews/gin-poc/controller"
	"gitlab.com/pragmaticreviews/gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/gin-poc/service"
)

// 사용할 변수를 묶어서 밖에 선언
var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// log 파일 생성
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	// 파일과 화면에 출력하기 위해 MultiWriter 생성
	// gin의 DefaultWriter로 지정(in mode.go)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	// server := gin.Default() 와 같음
	server := gin.New()
	server.Use(gin.Recovery(),
		middlewares.Logger(),    // custom logger
		middlewares.BasicAuth(), // custom auth
		gindump.Dump())          // header/body dump

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
