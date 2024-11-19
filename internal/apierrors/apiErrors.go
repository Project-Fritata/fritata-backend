package apierrors

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Status  string `json:"status" example:"<status code>: <status message>"`
	Message string `json:"message" validate:"omitempty" example:"<custom error message>"`
}

func DefaultError() error {
	return fmt.Errorf("something went wrong")
}

func getErrorResponseStatus(c fiber.Ctx) string {
	return fmt.Sprintf("%d: %s", c.Response().StatusCode(), http.StatusText(c.Response().StatusCode()))
}

func getHttpError(c fiber.Ctx, httpErrorCode int, err error) error {
	c.Status(httpErrorCode)

	errRes := ErrorResponse{
		Status:  getErrorResponseStatus(c),
		Message: err.Error(),
	}

	return c.JSON(errRes)
}

func InternalServerError(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusInternalServerError, err)
}
func InvalidRequest(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusBadRequest, err)
}
func InvalidCredentials(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusBadRequest, err)
}
func UserNotFound(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusNotFound, err)
}
func Unauthenticated(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusUnauthorized, err)
}
func UnprocessableEntity(c fiber.Ctx, err error) error {
	return getHttpError(c, fiber.StatusUnprocessableEntity, err)
}
