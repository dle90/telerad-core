package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"
	"telerad-core-module/utils"

	staffAccountControllerRequests "telerad-core-module/internals/requests/staff-account-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func StaffGetPaginatedStaffAccounts(c *fiber.Ctx) error {
	logger.Info("StaffGetPaginatedStaffAccounts starting....")

	page, size := utils.GetPaginationParams(c)
	search := utils.GetStringFromRequestParam(c, "search")

	isActive, err := utils.GetBoolFromRequestParam(c, "isActive", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetPaginatedStaffAccounts(c.Context(), page, size, search, isActive)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffGetAStaffAccount(c *fiber.Ctx) error {
	logger.Info("StaffGetAStaffAccount starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	result, systemErr := services.StaffGetAStaffAccount(c.Context(), staffUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffCreateStaffAccount(c *fiber.Ctx) error {
	logger.Info("StaffCreateStaffAccount starting....")

	var request staffAccountControllerRequests.StaffCreateStaffAccountRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	creatorUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCreateStaffAccount(c.Context(), creatorUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffUpdateStaffAccount(c *fiber.Ctx) error {
	logger.Info("StaffUpdateStaffAccount starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request staffAccountControllerRequests.StaffUpdateStaffAccountRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffUpdateStaffAccount(c.Context(), updaterUuid, staffUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffAssignReadingPermission(c *fiber.Ctx) error {
	logger.Info("StaffAssignReadingPermission starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request staffAccountControllerRequests.StaffAssignReadingPermissionRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffAssignReadingPermission(c.Context(), updaterUuid, staffUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffAssignRoles(c *fiber.Ctx) error {
	logger.Info("StaffAssignRoles starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request staffAccountControllerRequests.StaffAssignRolesRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffAssignRoles(c.Context(), updaterUuid, staffUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffActivateStaffAccount(c *fiber.Ctx) error {
	logger.Info("StaffActivateStaffAccount starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffActivateStaffAccount(c.Context(), updaterUuid, staffUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffDeactivateStaffAccount(c *fiber.Ctx) error {
	logger.Info("StaffDeactivateStaffAccount starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.StaffDeactivateStaffAccount(c.Context(), updaterUuid, staffUuid); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}

func StaffCreateAccount(c *fiber.Ctx) error {
	logger.Info("StaffCreateAccount starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request staffAccountControllerRequests.StaffCreateAccountRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCreateAccount(c.Context(), updaterUuid, staffUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func StaffResetStaffAccountPassword(c *fiber.Ctx) error {
	logger.Info("StaffResetStaffAccountPassword starting....")

	staffUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	updaterUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffResetPassword(c.Context(), updaterUuid, staffUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func UserGetMe(c *fiber.Ctx) error {
	logger.Info("UserGetMe starting....")

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.UserGetMe(c.Context(), userUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

func UserChangePassword(c *fiber.Ctx) error {
	logger.Info("UserChangePassword starting....")

	var request staffAccountControllerRequests.UserChangePasswordRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	if systemErr := services.UserChangePassword(c.Context(), userUuid, request); systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", nil))
}
