package imagingResultSheetTemplateControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetAImagingResultSheetTemplateResponse struct {
	Uuid               uuid.UUID  `json:"uuid"`
	CreatedAt          time.Time  `json:"createdAt"`
	CreatedBy          uuid.UUID  `json:"createdBy"`
	UpdatedAt          *time.Time `json:"updatedAt"`
	UpdatedBy          *uuid.UUID `json:"updatedBy"`
	TeleradPartnerUuid uuid.UUID  `json:"teleradPartnerUuid"`
	HtmlContent        string     `json:"htmlContent"`
	ResultFontSize     int16      `json:"resultFontSize"`
	ResultLineSpacing  float64    `json:"resultLineSpacing"`
	IsActive           bool       `json:"isActive"`
	//
	TeleradPartnerCode *string `json:"teleradPartnerCode"`
	TeleradPartnerName *string `json:"teleradPartnerName"`
}
