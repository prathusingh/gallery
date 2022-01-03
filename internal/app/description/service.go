package image

import "context"


type Image struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	GetImages(ctx context.Context)([]Image, error)
	
}

type imageService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &imageService{repository}
}

func (imgServ *imageService) GetImages(ctx context.Context) ([]Image) {
	return imgServ.repository.GetImages(ctx)
}