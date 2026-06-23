package services

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	fieldValues "telerad-core-module/internals/entities/field-values"
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
	status string, resultReturned *bool,
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
		Status:           strings.TrimSpace(status),
		ResultReturned:   resultReturned,
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

// StaffGetReadingOrderDetail trả chi tiết 1 ca đọc cho tab chi tiết. Áp dụng cùng
// kiểm tra quyền như khi mở viewer: user thường chỉ xem ca thuộc đối tác + loại chụp
// được phân; ADMIN không giới hạn.
func StaffGetReadingOrderDetail(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrderUuid uuid.UUID,
) (*teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse, *_error.SystemError) {
	readingOrder, systemErr := loadReadingOrderForStaff(ctx, requesterUuid, readingOrderUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	return buildReadingOrderDetailResponse(ctx, requesterUuid, readingOrder)
}

// StaffReceiveReadingOrder — "Nhận ca": user nhận 1 ca đang CHƯA ĐỌC để đọc.
// Điều kiện: ca đang UNREAD; user không đang đọc dở ca nào khác (mỗi user 1 ca READING).
// Tác động: status=READING, assigned_at=now, assigned_to=user. Trả về chi tiết ca.
func StaffReceiveReadingOrder(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrderUuid uuid.UUID,
) (*teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse, *_error.SystemError) {
	readingOrder, systemErr := loadReadingOrderForStaff(ctx, requesterUuid, readingOrderUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	// Điều kiện: ca phải đang CHƯA ĐỌC.
	if readingOrder.Status != fieldValues.TELERAD_READING_ORDER_STATUS_UNREAD.Value {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_003)
	}

	// Mỗi user chỉ đọc 1 ca tại 1 thời điểm: chặn nếu đang có ca READING khác.
	inProgress, err := repositories.FindOneReadingOrderByAssigneeAndStatus(
		ctx, bunNoTransaction, requesterUuid, fieldValues.TELERAD_READING_ORDER_STATUS_READING.Value,
	)
	if err != nil {
		return nil, _error.New(err)
	} else if inProgress != nil {
		return nil, _error.NewErrorByString(fmt.Sprintf(_errorMessages.TELERAD_E103_005, inProgress.OrderItemCode, inProgress.FullName))
	}

	now := time.Now()
	readingOrder.Status = fieldValues.TELERAD_READING_ORDER_STATUS_READING.Value
	readingOrder.AssignedAt = &now
	readingOrder.AssignedTo = &requesterUuid

	if err := baseServices.UpdateWholeTeleradReadingOrderRecord(ctx, bunNoTransaction, requesterUuid, readingOrder); err != nil {
		return nil, _error.New(err)
	}

	return buildReadingOrderDetailResponse(ctx, requesterUuid, readingOrder)
}

// StaffCancelReadingOrderLock — "Hủy khóa": nhả ca đang đọc về CHƯA ĐỌC.
// Điều kiện: ca đang READING và assigned_to = user. Tác động: status=UNREAD,
// assigned_at=null, assigned_to=null. Trả về chi tiết ca.
func StaffCancelReadingOrderLock(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrderUuid uuid.UUID,
) (*teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse, *_error.SystemError) {
	readingOrder, systemErr := loadReadingOrderForStaff(ctx, requesterUuid, readingOrderUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	// Điều kiện: ca phải đang ĐANG ĐỌC và do chính user này nhận.
	if readingOrder.Status != fieldValues.TELERAD_READING_ORDER_STATUS_READING.Value ||
		readingOrder.AssignedTo == nil || *readingOrder.AssignedTo != requesterUuid {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_004)
	}

	readingOrder.Status = fieldValues.TELERAD_READING_ORDER_STATUS_UNREAD.Value
	readingOrder.AssignedAt = nil
	readingOrder.AssignedTo = nil

	if err := baseServices.UpdateWholeTeleradReadingOrderRecord(ctx, bunNoTransaction, requesterUuid, readingOrder); err != nil {
		return nil, _error.New(err)
	}

	return buildReadingOrderDetailResponse(ctx, requesterUuid, readingOrder)
}

// StaffSaveReadingOrderResult — "Lưu kết quả": ghi nội dung kết quả (html) vào ca.
// Điều kiện: ca đang READING và assigned_to = user. Trả về chi tiết ca.
func StaffSaveReadingOrderResult(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrderUuid uuid.UUID,
	resultInHtml string,
) (*teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse, *_error.SystemError) {
	readingOrder, systemErr := loadReadingOrderForStaff(ctx, requesterUuid, readingOrderUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	if readingOrder.Status != fieldValues.TELERAD_READING_ORDER_STATUS_READING.Value ||
		readingOrder.AssignedTo == nil || *readingOrder.AssignedTo != requesterUuid {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_004)
	}

	readingOrder.ResultInHtml = &resultInHtml

	if err := baseServices.UpdateWholeTeleradReadingOrderRecord(ctx, bunNoTransaction, requesterUuid, readingOrder); err != nil {
		return nil, _error.New(err)
	}

	return buildReadingOrderDetailResponse(ctx, requesterUuid, readingOrder)
}

// loadReadingOrderForStaff tìm ca đọc + kiểm tra quyền truy cập (giống mở viewer):
// user thường chỉ thao tác ca thuộc đối tác + loại chụp được phân; ADMIN không giới hạn.
func loadReadingOrderForStaff(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrderUuid uuid.UUID,
) (*entities.TeleradReadingOrderEntity, *_error.SystemError) {
	readingOrder, err := baseServices.FindOneTeleradReadingOrderByUuid(ctx, bunNoTransaction, readingOrderUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if readingOrder == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_001)
	}

	staff, isAdmin, systemErr := resolveReadingScope(ctx, bunNoTransaction, requesterUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	if !isAdmin {
		if !slices.Contains(staff.TeleradPartnerUuids, readingOrder.TeleradPartnerUuid) {
			return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_002)
		} else if readingOrder.Modality == nil || !slices.Contains(staff.Modalities, *readingOrder.Modality) {
			return nil, _error.NewErrorByString(_errorMessages.TELERAD_E103_002)
		}
	}

	return readingOrder, nil
}

// buildReadingOrderDetailResponse resolve tên đối tác + tên bác sĩ đọc + cờ
// assignedToMe (ca có đang do chính requester đọc không) rồi map sang response.
func buildReadingOrderDetailResponse(
	ctx context.Context,
	requesterUuid uuid.UUID,
	readingOrder *entities.TeleradReadingOrderEntity,
) (*teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse, *_error.SystemError) {
	// Tên đối tác.
	partnerName := ""
	if partner, err := baseServices.FindOneTeleradPartnerByUuid(ctx, bunNoTransaction, readingOrder.TeleradPartnerUuid); err != nil {
		return nil, _error.New(err)
	} else if partner != nil {
		partnerName = partner.Name
	}

	// Tên bác sĩ đọc (nếu đã phân công).
	var assignedToName *string
	if readingOrder.AssignedTo != nil {
		if doctor, err := baseServices.FindOneStaffAccountByUuid(ctx, bunNoTransaction, *readingOrder.AssignedTo); err != nil {
			return nil, _error.New(err)
		} else if doctor != nil {
			assignedToName = &doctor.FullName
		}
	}

	assignedToMe := readingOrder.AssignedTo != nil && *readingOrder.AssignedTo == requesterUuid

	response := objectMappers.ToStaffGetReadingOrderDetailResponse(*readingOrder, partnerName, assignedToName, assignedToMe)
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
