package content

import "context"

// Image is used for displaying images
type Image struct {
	ImageID string `json:"id"`
	ImageName string `json:"name"`
}

// Description is used to describe the image
type Description struct {
	DescriptionID string `json:"id"`
	DescriptionName string `json:"name"`
}

type content struct {
	image Image
	description Description
}

type service interface {
	getContent(ctx context.Context, imageID string, descriptionID string)(content, error)
}

type contentService struct {
	repository Repository
}

func NewService(repository Repository) service {
	return &contentService{repository}
}

func (contentServ *contentService) getContent(ctx context.Context) (content, error) {
	return contentServ.repository.GetImages(ctx)
}