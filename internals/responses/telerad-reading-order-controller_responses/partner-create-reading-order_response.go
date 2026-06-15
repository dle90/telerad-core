package teleradReadingOrderControllerResponses

import "github.com/google/uuid"

// PartnerCreateReadingOrderResponse — xác nhận telerad đã nhận ca đọc.
type PartnerCreateReadingOrderResponse struct {
	Uuid uuid.UUID `json:"uuid"`
}
