package services

import (
	"context"
	"slices"
	"strings"
	"time"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	teleradReadingOrderControllerRequests "telerad-core-module/internals/requests/telerad-reading-order-controller_requests"
	"telerad-core-module/internals/responses"
	teleradReadingOrderControllerResponses "telerad-core-module/internals/responses/telerad-reading-order-controller_responses"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// PartnerCreateReadingOrder nhận 1 ca đọc do đối tác đẩy sang. teleradPartnerUuid lấy
// từ JWT — đồng thời dùng để chặn token không phải đối tác (uuid không khớp partner) và
// đối tác đã bị khóa. Trùng (partner + orderItemId) thì ghi đè thông tin lên bản ghi cũ
// (giữ nguyên trạng thái + phân công).
func PartnerCreateReadingOrder(
	ctx context.Context,
	teleradPartnerUuid uuid.UUID,
	request teleradReadingOrderControllerRequests.PartnerCreateReadingOrderRequest,
) (*teleradReadingOrderControllerResponses.PartnerCreateReadingOrderResponse, *_error.SystemError) {
	partner, err := baseServices.FindOneTeleradPartnerByUuid(ctx, bunNoTransaction, teleradPartnerUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if partner == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E101_001)
	} else if !partner.IsActive {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E001_002)
	}

	existing, err := repositories.FindOneReadingOrderByPartnerAndOrderItemId(ctx, bunNoTransaction, teleradPartnerUuid, request.OrderItemId)
	if err != nil {
		return nil, _error.New(err)
	}

	if existing != nil {
		baseServices.OverwriteTeleradReadingOrderInfo(existing, request)
		if err := baseServices.UpdateWholeTeleradReadingOrderRecord(ctx, bunNoTransaction, teleradPartnerUuid, existing); err != nil {
			return nil, _error.New(err)
		}

		response := objectMappers.ToPartnerCreateReadingOrderResponse(*existing)
		return &response, nil
	}

	readingOrder := baseServices.InitNewTeleradReadingOrder(teleradPartnerUuid, request)
	if err := baseServices.CreateNewTeleradReadingOrder(ctx, bunNoTransaction, teleradPartnerUuid, &readingOrder); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToPartnerCreateReadingOrderResponse(readingOrder)
	return &response, nil
}

// StaffGetPaginatedReadingOrders trả danh sách ca đọc cho màn chính, đã scope theo
// quyền user + lọc theo lựa chọn cây bên trái (partner/modality) và bộ lọc
// (ngày chụp, tên/mã bệnh nhân, số điện thoại).
func StaffGetPaginatedReadingOrders(
	ctx context.Context,
	userUuid uuid.UUID,
	page, pageSize int,
	selectedPartnerUuid *uuid.UUID,
	selectedModality string,
	performEndedFrom, performEndedTo *time.Time,
	patientName, patientCode, phone string,
) (*responses.PaginationResponse, *_error.SystemError) {
	staff, isAdmin, systemErr := resolveReadingScope(ctx, bunNoTransaction, userUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	filter := repositories.ReadingOrderListFilter{
		IsAdmin:          isAdmin,
		PartnerUuids:     staff.TeleradPartnerUuids,
		Modalities:       staff.Modalities,
		PerformEndedFrom: performEndedFrom,
		PerformEndedTo:   performEndedTo,
		PatientName:      strings.TrimSpace(patientName),
		PatientCode:      strings.TrimSpace(patientCode),
		Phone:            strings.TrimSpace(phone),
	}

	if isAdmin {
		if selectedPartnerUuid != nil {
			filter.PartnerUuids = []uuid.UUID{*selectedPartnerUuid}
		} else {
			filter.PartnerUuids = nil // admin không chọn partner nào → xem tất cả partner
		}

		if selectedModality != "" {
			filter.Modalities = []string{selectedModality}
		} else {
			filter.Modalities = nil // admin không chọn modality nào → xem tất cả modality
		}
	} else {
		if selectedPartnerUuid != nil {
			// User thường chỉ được chọn partner trong phạm vi được phân. Nếu chọn partner ngoài phạm vi → không có ca nào để xem.
			if !slices.Contains(staff.TeleradPartnerUuids, *selectedPartnerUuid) {
				filter.PartnerUuids = nil
			} else {
				filter.PartnerUuids = []uuid.UUID{*selectedPartnerUuid}
			}
		}

		if selectedModality != "" {
			// User thường chỉ được chọn modality trong phạm vi được phân. Nếu chọn modality ngoài phạm vi → không có ca nào để xem.
			if !slices.Contains(staff.Modalities, selectedModality) {
				filter.Modalities = nil
			} else {
				filter.Modalities = []string{selectedModality}
			}
		}
	}

	rows, totalCount, err := repositories.FindPaginatedReadingOrders(ctx, bunNoTransaction, page, pageSize, filter)
	if err != nil {
		return nil, _error.New(err)
	}

	records := objectMappers.ToStaffGetListReadingOrderSlice(rows)
	response := responses.NewPaginationResponse(totalCount, page, pageSize, records)
	return &response, nil
}

// resolveReadingScope tìm staff đang đăng nhập và xác định có phải ADMIN không.
// ADMIN không bị giới hạn phạm vi đọc; user thường bị giới hạn theo modalities +
// telerad_partner_uuids gắn trên hồ sơ.
func resolveReadingScope(ctx context.Context, tx bun.IDB, userUuid uuid.UUID) (*entities.StaffAccountEntity, bool, *_error.SystemError) {
	staff, err := baseServices.FindOneStaffAccountByUuid(ctx, tx, userUuid)
	if err != nil {
		return nil, false, _error.New(err)
	} else if staff == nil {
		return nil, false, _error.NewErrorByString(_errorMessages.TELERAD_E102_001)
	}

	return staff, slices.Contains(staff.Roles, constants.ROLE_ADMIN), nil
}
