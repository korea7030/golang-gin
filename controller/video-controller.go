package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/pragmaticreviews/gin-poc/entity"
	"gitlab.com/pragmaticreviews/gin-poc/service"
	"gitlab.com/pragmaticreviews/gin-poc/validators"
)

// controller interface: 사용할 함수를 선언
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

// controller struct : service와 연결 역할
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	// validation 등록(is-cool 이란 이름으로 ValidateCoolTitle 함수를 등록하겠다)
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

// VideoController interface의 함수 선언
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// VideoController interface의 함수 선언
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	// binding 사용하기 위해 변경
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)

	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil

}

// VideoController interface의 html load 함수
func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	video.ID = id
	err = validate.Struct(video)

	if err != nil {
		return nil
	}

	c.service.Update(video)
	return nil

}
func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}

	video.ID = id
	c.service.Delete(video)
	return nil
}
