package pkg

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-restful-fiber/model/dto"
)

func NewErrorHandler(ctx *fiber.Ctx, err error) error {
	// Init Logger
	logger := NewLogger()

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return if Not Found
	if code == fiber.StatusNotFound {
		return ctx.Status(code).JSON(dto.ApiResponseFail{
			Success: false,
			Message: "Not Found",
		})
	}

	// Return if Bad Request
	if code == fiber.StatusBadRequest {
		return ctx.Status(code).JSON(dto.ApiResponseError{
			Success:   false,
			Message:   "Bad Request",
			ErrorCode: "BAD_REQUEST",
			Errors:    err.Error(),
		})
	}

	// Return if Unauthorized
	if code == fiber.StatusUnauthorized {
		return ctx.Status(code).JSON(dto.ApiResponseFail{
			Success: false,
			Message: "Unauthorized",
		})
	}

	// Return if Forbidden
	if code == fiber.StatusForbidden {
		return ctx.Status(code).JSON(dto.ApiResponseFail{
			Success: false,
			Message: "Forbidden",
		})
	}

	// Logging Error
	logger.Error(err)

	// Return Internal Server Error
	return ctx.Status(code).JSON(dto.ApiResponseFail{
		Success: false,
		Message: "Internal Server Error",
	})
}
