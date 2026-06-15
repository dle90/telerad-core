package controllers

import (
	"net/http"

	authControllerRequests "telerad-core-module/internals/requests/auth-controller_requests"
	"telerad-core-module/internals/services"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func StaffLogin(c *fiber.Ctx) error {
	logger.Info("StaffLogin starting....")

	var request authControllerRequests.StaffLoginRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffLogin(c.Context(), request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func TeleradPartnerLogin(c *fiber.Ctx) error {
	logger.Info("TeleradPartnerLogin starting....")

	var request authControllerRequests.TeleradPartnerLoginRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.TeleradPartnerLogin(c.Context(), request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(result)
}
