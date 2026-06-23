package imagingResultSheetTemplateControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetListImagingResultSheetTemplateResponse struct {
	Uuid               uuid.UUID `json:"uuid"`
	TeleradPartnerUuid uuid.UUID `json:"teleradPartnerUuid"`
	TeleradPartnerCode *string   `json:"teleradPartnerCode"`
	TeleradPartnerName *string   `json:"teleradPartnerName"`
	IsActive           bool      `json:"isActive"`
	CreatedAt          time.Time `json:"createdAt"`
}
