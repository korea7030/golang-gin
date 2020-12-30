package service

// LoginService interface
type LoginService interface {
	Login(username string, password string) bool
}

// authorization LoginService
type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// loginService New
func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "pragmatic",
		authorizedPassword: "reviews",
	}
}

// LoginService interface 함수
func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username && service.authorizedPassword == password
}
