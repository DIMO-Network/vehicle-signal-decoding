package common

import (
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig(disableStartupMsg bool) fiber.Config {
	return fiber.Config{
		DisableStartupMessage: disableStartupMsg,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError
			_type := "https://tools.ietf.org/html/rfc7231#section-6.6.1"
			title := "An error occurred while processing your request."

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if _, ok := err.(*exceptions.ValidationError); ok {
				code = fiber.StatusBadRequest
				_type = "https://tools.ietf.org/html/rfc7231#section-6.5.1"
				title = "The specified resource is not valid."
			}

			if _, ok := err.(*exceptions.NotFoundError); ok {
				code = fiber.StatusNotFound
				_type = "https://tools.ietf.org/html/rfc7231#section-6.5.4"
				title = "The specified resource was not found."
			}

			if _, ok := err.(*exceptions.ConflictError); ok {
				code = fiber.StatusConflict
				_type = "https://tools.ietf.org/html/rfc7231#section-6.5.1"
				title = "The specified resource is not valid."
			}

			p := &ProblemDetails{
				Type:   _type,
				Title:  title,
				Status: code,
				Detail: err.Error(),
			}

			return ctx.Status(code).JSON(p)
		},
	}
}
