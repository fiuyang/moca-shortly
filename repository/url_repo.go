package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"scylla/model"
	"scylla/pkg/helper"
)

type UrlRepo interface {
	Insert(ctx context.Context, data model.URL) (model.URL, error)
	FindByShortCode(ctx context.Context, shortCode string) (data model.URL, err error)
	IncrementClicks(ctx context.Context, data model.URL) error
}

type UrlRepoImpl struct {
	db *gorm.DB
}

func NewUrlRepoImpl(db *gorm.DB) UrlRepo {
	return &UrlRepoImpl{
		db: db,
	}
}

func (repo *UrlRepoImpl) Insert(ctx context.Context, data model.URL) (model.URL, error) {
	result := repo.db.WithContext(ctx).Create(&data)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
	}
	return data, nil
}

func (repo *UrlRepoImpl) FindByShortCode(ctx context.Context, shortCode string) (data model.URL, err error) {
	result := repo.db.WithContext(ctx).Where("short_code", shortCode).First(&data)
	if result.RowsAffected == 0 {
		return data, errors.New("record not found")
	}
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

func (repo *UrlRepoImpl) IncrementClicks(ctx context.Context, data model.URL) error {
	data.Clicks++
	return repo.db.WithContext(ctx).Save(data).Error
}
