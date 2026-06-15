package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"
	"telerad-core-module/utils"

	teleradReadingOrderControllerRequests "telerad-core-module/internals/requests/telerad-reading-order-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func PartnerCreateReadingOrder(c *fiber.Ctx) error {
	logger.Info("PartnerCreateReadingOrder starting....")

	var request teleradReadingOrderControllerRequests.PartnerCreateReadingOrderRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	teleradPartnerUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.PartnerCreateReadingOrder(c.Context(), teleradPartnerUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffGetTeleradPartnersForReading — cây bên trái màn "Đọc ca": đối tác nhóm
// theo loại chụp, lọc theo quyền user (ADMIN xem tất cả).
func StaffGetTeleradPartnersForReading(c *fiber.Ctx) error {
	logger.Info("StaffGetTeleradPartnersForReading starting....")

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffGetTeleradPartnersForReading(c.Context(), userUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffGetPaginatedReadingOrders — danh sách ca đọc màn chính, scope theo quyền user
// + lọc theo partner/modality đang chọn và ngày chụp / tên / mã bệnh nhân / SĐT.
func StaffGetPaginatedReadingOrders(c *fiber.Ctx) error {
	logger.Info("StaffGetPaginatedReadingOrders starting....")

	page, size := utils.GetPaginationParams(c)

	selectedPartnerUuid, err := utils.GetUuidFromRequestParam(c, "teleradPartnerUuid", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	selectedModality := utils.GetStringFromRequestParam(c, "modality")
	patientName := utils.GetStringFromRequestParam(c, "patientName")
	patientCode := utils.GetStringFromRequestParam(c, "patientCode")
	phone := utils.GetStringFromRequestParam(c, "phone")

	performEndedFrom, err := utils.GetTimeFromRequestParam(c, "performEndedFrom", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	performEndedTo, err := utils.GetTimeFromRequestParam(c, "performEndedTo", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffGetPaginatedReadingOrders(
		c.Context(), userUuid, page, size,
		selectedPartnerUuid, selectedModality,
		performEndedFrom, performEndedTo,
		patientName, patientCode, phone,
	)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffGenerateImagingStudyViewerUrl — sinh URL mở PACS viewer cho 1 ca đọc (kèm
// view-token trong URL hash). Chỉ cho mở ca thuộc quyền của user (service kiểm tra).
func StaffGenerateImagingStudyViewerUrl(c *fiber.Ctx) error {
	logger.Info("StaffGenerateImagingStudyViewerUrl starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	requesterUuid := secure.GetUserUuidFromJwt(c)
	requesterUsername := secure.GetUsernameFromJwt(c)

	result, systemErr := services.StaffGenerateImagingStudyViewerUrl(c.Context(), requesterUuid, requesterUsername, readingOrderUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}
