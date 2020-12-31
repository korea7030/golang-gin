package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	"gitlab.com/pragmaticreviews/gin-poc/api"
	"gitlab.com/pragmaticreviews/gin-poc/controller"
	"gitlab.com/pragmaticreviews/gin-poc/docs" // Swagger generated files
	"gitlab.com/pragmaticreviews/gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/gin-poc/repository"
	"gitlab.com/pragmaticreviews/gin-poc/service"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// 사용할 변수를 묶어서 밖에 선언
var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// log 파일 생성
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	// 파일과 화면에 출력하기 위해 MultiWriter 생성
	// gin의 DefaultWriter로 지정(in mode.go)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization

func main() {

	docs.SwaggerInfo.Title = "Progmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic ReViews - Youtube Video API"
	docs.SwaggerInfo.Host = "pragmatic-video-app.herokuapp.com"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	setupLogOutput()

	defer videoRepository.CloseDB()
	// server := gin.Default() 와 같음
	defer videoRepository.CloseDB()

	server := gin.Default()

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	// swagger url로 요청시 다음 호출
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "7070"
	}
	server.Run(":" + port)

}
