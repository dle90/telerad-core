package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"
	"telerad-core-module/utils"

	imagingResultSheetTemplateControllerRequests "telerad-core-module/internals/requests/imaging-result-sheet-template-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func StaffGetPaginatedImagingResultSheetTemplates(c *fiber.Ctx) error {
	logger.Info("StaffGetPaginatedImagingResultSheetTemplates starting....")

	page, size := utils.GetPaginationParams(c)

	teleradPartnerUuid, err := utils.GetUuidFromRequestParam(c, "teleradPartnerUuid", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	isActive, err := utils.GetBoolFromRequestParam(c, "isActive", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetPaginatedImagingResultSheetTemplates(c.Context(), page, size, teleradPartnerUuid, isActive)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetAImagingResultSheetTemplate(c *fiber.Ctx) error {
	logger.Info("StaffGetAImagingResultSheetTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetAImagingResultSheetTemplate(c.Context(), templateUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffCreateImagingResultSheetTemplate(c *fiber.Ctx) error {
	logger.Info("StaffCreateImagingResultSheetTemplate starting....")

	var request imagingResultSheetTemplateControllerRequests.StaffCreateImagingResultSheetTemplateRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	creatorUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCreateImagingResultSheetTemplate(c.Context(), creatorUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffUpdateImagingResultSheetTemplate(c *fiber.Ctx) error {
	logger.Info("StaffUpdateImagingResultSheetTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request imagingResultSheetTemplateControllerRequests.StaffUpdateImagingResultSheetTemplateRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffUpdateImagingResultSheetTemplate(c.Context(), updaterUuid, templateUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffActivateImagingResultSheetTemplate(c *fiber.Ctx) error {
	logger.Info("StaffActivateImagingResultSheetTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffActivateImagingResultSheetTemplate(c.Context(), updaterUuid, templateUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffDeactivateImagingResultSheetTemplate(c *fiber.Ctx) error {
	logger.Info("StaffDeactivateImagingResultSheetTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffDeactivateImagingResultSheetTemplate(c.Context(), updaterUuid, templateUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}
