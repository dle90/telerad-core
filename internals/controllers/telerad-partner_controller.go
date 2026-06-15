package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"
	"telerad-core-module/utils"

	teleradPartnerControllerRequests "telerad-core-module/internals/requests/telerad-partner-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func StaffGetPaginatedTeleradPartners(c *fiber.Ctx) error {
	logger.Info("StaffGetPaginatedTeleradPartners starting....")

	page, size := utils.GetPaginationParams(c)
	search := utils.GetStringFromRequestParam(c, "search")

	isActive, err := utils.GetBoolFromRequestParam(c, "isActive", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetPaginatedTeleradPartners(c.Context(), page, size, search, isActive)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetAllTeleradPartners(c *fiber.Ctx) error {
	logger.Info("StaffGetAllTeleradPartners starting....")

	result, systemErr := services.StaffGetAllTeleradPartners(c.Context())
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetATeleradPartner(c *fiber.Ctx) error {
	logger.Info("StaffGetATeleradPartner starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetATeleradPartner(c.Context(), partnerUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffCreateTeleradPartner(c *fiber.Ctx) error {
	logger.Info("StaffCreateTeleradPartner starting....")

	var request teleradPartnerControllerRequests.StaffCreateTeleradPartnerRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	creatorUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCreateTeleradPartner(c.Context(), creatorUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffUpdateTeleradPartner(c *fiber.Ctx) error {
	logger.Info("StaffUpdateTeleradPartner starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request teleradPartnerControllerRequests.StaffUpdateTeleradPartnerRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffUpdateTeleradPartner(c.Context(), updaterUuid, partnerUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffActivateTeleradPartner(c *fiber.Ctx) error {
	logger.Info("StaffActivateTeleradPartner starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffActivateTeleradPartner(c.Context(), updaterUuid, partnerUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffDeactivateTeleradPartner(c *fiber.Ctx) error {
	logger.Info("StaffDeactivateTeleradPartner starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffDeactivateTeleradPartner(c.Context(), updaterUuid, partnerUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffGetTeleradPartnerPartnerConfig(c *fiber.Ctx) error {
	logger.Info("StaffGetTeleradPartnerPartnerConfig starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetTeleradPartnerPartnerConfig(c.Context(), partnerUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffUpdateTeleradPartnerPartnerConfig(c *fiber.Ctx) error {
	logger.Info("StaffUpdateTeleradPartnerPartnerConfig starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request teleradPartnerControllerRequests.StaffUpdateTeleradPartnerPartnerConfigRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffUpdateTeleradPartnerPartnerConfig(c.Context(), updaterUuid, partnerUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffChangeTeleradPartnerPassword(c *fiber.Ctx) error {
	logger.Info("StaffChangeTeleradPartnerPassword starting....")

	partnerUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request teleradPartnerControllerRequests.StaffChangeTeleradPartnerPasswordRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffChangeTeleradPartnerPassword(c.Context(), updaterUuid, partnerUuid, request); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}
