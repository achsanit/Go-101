package repository

import (
	"context"
	"errors"

	"github.com/achsanit/my-gram/internal/infrastructure"
	"github.com/achsanit/my-gram/internal/model"
)

type PhotoQuery interface {
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	GetPhotosUser(ctx context.Context, userId int) ([]model.PhotoUser, error)

	GetPhotoById(ctx context.Context, id int) (model.Photo, error)
}

type photoQueryImpl struct {
	db infrastructure.GormPostgres
}

// GetPhotoById implements PhotoQuery.
func (p *photoQueryImpl) GetPhotoById(ctx context.Context, id int) (model.Photo, error) {
	db := p.db.GetConnection()

	photo := model.Photo{}
	result := db.WithContext(ctx).Where("id = ?", id).First(&photo)
	if result.Error != nil {
		return model.Photo{}, errors.New("photo not found")
	}

	return photo, nil
}

// NEED IMPROVEMENT....
// GetPhotosUser implements PhotoQuery.
func (p *photoQueryImpl) GetPhotosUser(ctx context.Context, userId int) ([]model.PhotoUser, error) {
	db := p.db.GetConnection()

	userPhotos := []model.PhotoUser{}
	if err := db.
		WithContext(ctx).
		Table("photos").
		Find(&userPhotos).Error; err != nil {
		return nil, err
	}

	return userPhotos, nil
}

// CreatePhoto implements PhotoQuery.
func (p *photoQueryImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := p.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("photos").
		Save(&photo).
		Error; err != nil {

		return model.Photo{}, err
	}

	return photo, nil
}

func NewPhotoQuery(db infrastructure.GormPostgres) PhotoQuery {
	return &photoQueryImpl{
		db: db,
	}
}
