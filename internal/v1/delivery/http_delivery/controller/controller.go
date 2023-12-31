package controller

import (
	"net/http"

	"github.com/dathuynh1108/clean-arch-base/internal/v1/entity"
	"github.com/dathuynh1108/clean-arch-base/pkg/comerr"
	"github.com/dathuynh1108/clean-arch-base/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type controller struct{}

func (c *controller) BindAndValidate(ctx *fiber.Ctx, data any) error {
	if err := ctx.BodyParser(data); err != nil {
		return comerr.WrapError(err, "Failed to parse request body")
	}
	errs := validator.GetValidator().Validate(data)
	if errs != nil {
		return c.Failure(ctx, http.StatusBadRequest, http.StatusBadRequest, "Invalid request", errs)
	}
	return nil
}

func (c *controller) OK(ctx *fiber.Ctx, code int, message any, data any) error {
	return ctx.
		Status(http.StatusOK).
		JSON(&entity.Response{
			Code:    code,
			Message: message,
			Data:    data,
		})
}

func (c *controller) OKEmpty(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(nil)
}

func (c *controller) Failure(ctx *fiber.Ctx, httpCode int, code int, message any, errors []error) error {
	return ctx.
		Status(httpCode).
		JSON(&entity.Response{
			Code:    code,
			Message: message,
			Data:    nil,
			Errors:  errors,
		})
}
