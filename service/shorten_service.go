package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"scylla/entity"
	"scylla/model"
	"scylla/pkg/exception"
	"scylla/pkg/helper"
	"scylla/pkg/utils"
	"scylla/repository"
	"time"
)

type ShortenService interface {
	ShortenURL(ctx context.Context, request entity.ShortenURLRequest) (response entity.ShortenURLResponse, err error)
	GetOriginalURL(ctx context.Context, shortCode string) (response entity.RedirectResponse, err error)
}

type ShortenServiceImpl struct {
	urlRepo  repository.UrlRepo
	validate *validator.Validate
}

func NewShortenServiceImpl(urlRepo repository.UrlRepo, validate *validator.Validate) ShortenService {
	return &ShortenServiceImpl{
		urlRepo:  urlRepo,
		validate: validate,
	}
}

func (service *ShortenServiceImpl) ShortenURL(ctx context.Context, request entity.ShortenURLRequest) (response entity.ShortenURLResponse, err error) {
	err = service.validate.Struct(request)
	helper.ErrorPanic(err)

	expiredAt, err := time.Parse("2006-01-02", request.Expiry) //yyyy-mm-dd
	helper.ErrorPanic(err)

	shortCode := utils.GenerateShortCode()

	dataset := model.URL{
		ShortCode:   shortCode,
		OriginalURL: request.OriginalURL,
		ExpiredAt:   &expiredAt,
	}

	result, err := service.urlRepo.Insert(ctx, dataset)
	if err != nil {
		panic(exception.NewInternalServerErrorHandler(err.Error()))
	}

	response.Expiry = result.ExpiredAt
	helper.Automapper(result, &response)
	return response, nil
}

func (service *ShortenServiceImpl) GetOriginalURL(ctx context.Context, shortCode string) (response entity.RedirectResponse, err error) {
	result, err := service.urlRepo.FindByShortCode(ctx, shortCode)
	if err != nil {
		panic(exception.NewNotFoundHandler(err.Error()))
	}

	if result.ExpiredAt != nil && time.Now().After(*result.ExpiredAt) {
		return response, errors.New("URL has expired")
	}

	err = service.urlRepo.IncrementClicks(ctx, result)
	if err != nil {
		return response, err
	}

	helper.Automapper(result, &response)
	return response, nil
}
