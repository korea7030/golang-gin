package main

import (
	"io"
	"net/http"
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

	// static 및 html load 방법
	server.Static("/css", "./templates/css") // 상대경로, root경로
	server.LoadHTMLGlob("templates/*.html")  // pattern 값으로 html 불러옴

	server.Use(gin.Recovery(),
		middlewares.Logger(), // custom logger
		gindump.Dump())       // header/body dump

	// api용 url 따로 분리(/api/videos) (authorization required)
	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		// video 정보 get
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		// video 정보 post
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input Is Valid!"})
			}
		})
	}

	// view용 url 따로 분리(/view/videos) (no Authorization required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "7070"
	}
	server.Run(":" + port)

}
