package errors

import (
	"telerad-core-module/internals/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handle lỗi hệ thống, luu y: phai chac chan "err != nil"
func HandleSystemError(c *fiber.Ctx, err *SystemError) error {
	if err == nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.NewBaseResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))
	} else if err.ErrorCode() == CHECKSUM_INVALID {
		return c.Status(http.StatusUnauthorized).JSON(responses.NewBaseResponse(http.StatusUnauthorized, err.ErrorMessage(), nil))
	} else if err.ErrorMessage() != "" {
		return c.Status(http.StatusBadRequest).JSON(responses.NewBaseResponse(http.StatusBadRequest, err.ErrorMessage(), nil))
	} else {
		return c.Status(http.StatusInternalServerError).JSON(responses.NewBaseResponse(http.StatusInternalServerError, err.ErrorMessage(), nil))
	}
}
