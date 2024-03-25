package service

import (
	"context"

	"github.com/achsanit/my-gram/internal/model"
	"github.com/achsanit/my-gram/internal/repository"
)

type PhotoService interface {
	PostPhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	GetPhotosUser(ctx context.Context, userId int) ([]model.PhotoUser, error)

	GetPhotoByID(ctx context.Context, id int) (model.Photo, error)
}

type photoServiceImpl struct {
	repo repository.PhotoQuery
}

// GetPhotoByID implements PhotoService.
func (p *photoServiceImpl) GetPhotoByID(ctx context.Context, id int) (model.Photo, error) {
	res, err := p.repo.GetPhotoById(ctx, id)
	if err != nil {
		return model.Photo{}, err
	}

	return res, nil
}

// GetPhotosUser implements PhotoService.
func (p *photoServiceImpl) GetPhotosUser(ctx context.Context, userId int) ([]model.PhotoUser, error) {
	res, err := p.repo.GetPhotosUser(ctx, userId)
	if err != nil {
		return []model.PhotoUser{}, err
	}

	return res, nil
}

// PostPhoto implements PhotoService.
func (p *photoServiceImpl) PostPhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	res, err := p.repo.CreatePhoto(ctx, photo)

	if err != nil {
		return model.Photo{}, err
	}

	return res, nil
}

func NewPhotoService(repo repository.PhotoQuery) PhotoService {
	return &photoServiceImpl{
		repo: repo,
	}
}
