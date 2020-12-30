package service

import (
	"gitlab.com/pragmaticreviews/gin-poc/entity"
	"gitlab.com/pragmaticreviews/gin-poc/repository"
)

// interface
type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

// DB repository 연결
type videoService struct {
	// videos []entity.Video
	videoRepository repository.VideoRepository
}

// New() : Service와 entity 연결
func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

// interface function
func (service *videoService) Save(video entity.Video) entity.Video {
	// service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

// interface function
func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

// interface function
func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

// interface function
func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
