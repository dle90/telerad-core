package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"
	"telerad-core-module/utils"

	imagingResultTemplateControllerRequests "telerad-core-module/internals/requests/imaging-result-template-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func StaffGetPaginatedImagingResultTemplates(c *fiber.Ctx) error {
	logger.Info("StaffGetPaginatedImagingResultTemplates starting....")

	page, size := utils.GetPaginationParams(c)
	search := utils.GetStringFromRequestParam(c, "search")
	modality := utils.GetStringFromRequestParam(c, "modality")

	isActive, err := utils.GetBoolFromRequestParam(c, "isActive", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	bodyParts := utils.GetStringSliceFromRequestParam(c, "bodyParts")

	result, systemErr := services.StaffGetPaginatedImagingResultTemplates(c.Context(), page, size, modality, search, isActive, bodyParts)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetImagingResultTemplateFormOptions(c *fiber.Ctx) error {
	logger.Info("StaffGetImagingResultTemplateFormOptions starting....")

	result := services.StaffGetImagingResultTemplateFormOptions()
	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetAImagingResultTemplate(c *fiber.Ctx) error {
	logger.Info("StaffGetAImagingResultTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetAImagingResultTemplate(c.Context(), templateUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffCreateImagingResultTemplate(c *fiber.Ctx) error {
	logger.Info("StaffCreateImagingResultTemplate starting....")

	var request imagingResultTemplateControllerRequests.StaffCreateImagingResultTemplateRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	creatorUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCreateImagingResultTemplate(c.Context(), creatorUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffUpdateImagingResultTemplate(c *fiber.Ctx) error {
	logger.Info("StaffUpdateImagingResultTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request imagingResultTemplateControllerRequests.StaffUpdateImagingResultTemplateRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffUpdateImagingResultTemplate(c.Context(), updaterUuid, templateUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffActivateImagingResultTemplate(c *fiber.Ctx) error {
	logger.Info("StaffActivateImagingResultTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffActivateImagingResultTemplate(c.Context(), updaterUuid, templateUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffDeactivateImagingResultTemplate(c *fiber.Ctx) error {
	logger.Info("StaffDeactivateImagingResultTemplate starting....")

	templateUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffDeactivateImagingResultTemplate(c.Context(), updaterUuid, templateUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}
