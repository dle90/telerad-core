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
	status := utils.GetStringFromRequestParam(c, "status")

	resultReturned, err := utils.GetBoolFromRequestParam(c, "resultReturned", true)
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

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
		status, resultReturned,
	)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffGetReadingOrderDetail — chi tiết 1 ca đọc cho tab chi tiết màn "Đọc ca".
func StaffGetReadingOrderDetail(c *fiber.Ctx) error {
	logger.Info("StaffGetReadingOrderDetail starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffGetReadingOrderDetail(c.Context(), userUuid, readingOrderUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffReceiveReadingOrder — "Nhận ca": user nhận 1 ca CHƯA ĐỌC để đọc.
func StaffReceiveReadingOrder(c *fiber.Ctx) error {
	logger.Info("StaffReceiveReadingOrder starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffReceiveReadingOrder(c.Context(), userUuid, readingOrderUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffCancelReadingOrderLock — "Hủy khóa": nhả ca đang đọc về CHƯA ĐỌC.
func StaffCancelReadingOrderLock(c *fiber.Ctx) error {
	logger.Info("StaffCancelReadingOrderLock starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffCancelReadingOrderLock(c.Context(), userUuid, readingOrderUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffSaveReadingOrderResult — "Lưu kết quả": ghi result_in_html cho ca đang đọc.
func StaffSaveReadingOrderResult(c *fiber.Ctx) error {
	logger.Info("StaffSaveReadingOrderResult starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	var request teleradReadingOrderControllerRequests.StaffSaveReadingOrderResultRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffSaveReadingOrderResult(c.Context(), userUuid, readingOrderUuid, request.ResultInHtml)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffEndReadingAndApprove — "Kết thúc & Duyệt": chốt ca đang đọc thành ĐÃ DUYỆT.
func StaffEndReadingAndApprove(c *fiber.Ctx) error {
	logger.Info("StaffEndReadingAndApprove starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffEndReadingAndApprove(c.Context(), userUuid, readingOrderUuid)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}

// StaffGetReadingOrderResultSheet — mẫu phiếu kết quả của CSYT ca đọc (để in).
func StaffGetReadingOrderResultSheet(c *fiber.Ctx) error {
	logger.Info("StaffGetReadingOrderResultSheet starting....")

	readingOrderUuid, err := utils.GetUuidFromRequestPath(c, "objectId")
	if err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	userUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.StaffGetReadingOrderResultSheet(c.Context(), userUuid, readingOrderUuid)
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
