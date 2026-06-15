package services

import (
	"context"

	baseServices "telerad-core-module/internals/base-services"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	teleradReadingOrderControllerRequests "telerad-core-module/internals/requests/telerad-reading-order-controller_requests"
	teleradReadingOrderControllerResponses "telerad-core-module/internals/responses/telerad-reading-order-controller_responses"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
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
