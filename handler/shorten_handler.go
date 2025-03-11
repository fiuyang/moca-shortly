package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"scylla/entity"
	"scylla/pkg/helper"
	"scylla/pkg/utils"
	"scylla/service"
	"time"
)

type ShortenHandler struct {
	shortenService service.ShortenService
}

func NewShortenHandler(shortenService service.ShortenService) *ShortenHandler {
	return &ShortenHandler{
		shortenService: shortenService,
	}
}

// Note            godoc
//
//	@Summary		Create Short URL
//	@Description	Create a short URL from an original URL.
//	@Param			data	body	entity.ShortenURLRequest		true	"create short URL"
//	@Produce		application/json
//	@Tags			shortener
//	@Success		201	{object}	entity.JsonCreated{data=entity.ShortenURLResponse{}}"Data"
//	@Failure		400	{object}	entity.JsonBadRequest{}				"Validation error"
//	@Failure		404	{object}	entity.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	entity.JsonInternalServerError{}	"Internal server error"
//	@Router			/shorten [post]
func (handler *ShortenHandler) ShortenURL(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := entity.ShortenURLRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	result, errs := handler.shortenService.ShortenURL(c, request)
	helper.ErrorPanic(errs)

	baseURL := ctx.BaseURL()
	shortURL := baseURL + "/api/v1/" + result.ShortCode

	response := entity.ShortenURLResponse{
		ShortCode: result.ShortCode,
		ShortURL:  shortURL,
		Expiry:    result.Expiry,
	}

	webResponse := entity.Response{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Short URL successfully created",
		Data:    response,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note 		    godoc
//
//	@Summary		Get Original URL.
//	@Param			shortCode	path	string	true	"Short Code"
//	@Description	Fetch the original URL from the short code.
//	@Produce		application/json
//	@Tags			shortener
//	@Success		200	{object}	entity.JsonSuccess{data=entity.RedirectResponse{}}	"Data"
//	@Failure		400	{object}	entity.JsonBadRequest{}								"Validation error"
//	@Failure		404	{object}	entity.JsonNotFound{}								"Data not found"
//	@Failure		500	{object}	entity.JsonInternalServerError{}					"Internal server error"
//	@Router			/shorten/{shortCode} [get]
func (handler *ShortenHandler) ShowOriginalURL(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	shortCode := ctx.Params("shortCode")

	data, err := handler.shortenService.GetOriginalURL(c, shortCode)
	helper.ErrorPanic(err)

	webResponse := entity.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "success",
		Data:    data,
	}

	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (handler *ShortenHandler) RedirectURL(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	shortCode := ctx.Params("shortCode")

	data, err := handler.shortenService.GetOriginalURL(c, shortCode)
	helper.ErrorPanic(err)

	return ctx.Redirect(data.OriginalURL, fiber.StatusFound)
}
