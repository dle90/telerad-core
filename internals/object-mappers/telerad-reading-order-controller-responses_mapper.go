package objectMappers

import (
	"telerad-core-module/internals/entities"
	teleradReadingOrderControllerResponses "telerad-core-module/internals/responses/telerad-reading-order-controller_responses"
)

func ToPartnerCreateReadingOrderResponse(readingOrder entities.TeleradReadingOrderEntity) teleradReadingOrderControllerResponses.PartnerCreateReadingOrderResponse {
	return teleradReadingOrderControllerResponses.PartnerCreateReadingOrderResponse{
		Uuid: readingOrder.Uuid,
	}
}
