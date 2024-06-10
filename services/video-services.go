package services

import (
	"gin-project/entities"
	"gin-project/repository"
)

type VideoService interface {
	Save(entities.Video) entities.Video
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() []entities.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entities.Video) entities.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entities.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entities.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []entities.Video {
	return service.videoRepository.FindAll()
}
