package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/gin-poc/dto"
	"gitlab.com/pragmaticreviews/gin-poc/service"
)

// LoginController interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

// loginService와 jwtService 연결
type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

// Controller 생성 시, loginService, jwtService 연결
func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

// logincontroller Login 함수
func (controller *loginController) Login(ctx *gin.Context) string {
	// credential bind
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)

	if err != nil {
		return ""
	}

	// loginService 호출
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	// true시 token generate
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
