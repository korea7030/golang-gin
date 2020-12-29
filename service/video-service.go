package service

import (
	"gitlab.com/pragmaticreviews/gin-poc/entity"
)

// interface
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// New() : Service와 entity 연결
func New() VideoService {
	return &videoService{}
}

// interface function
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// interface function
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
