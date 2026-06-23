package imagingResultSheetTemplateControllerRequests

import "github.com/google/uuid"

// StaffCreateImagingResultSheetTemplateRequest — tạo mẫu phiếu kết quả cho 1 CSYT
// (telerad_partner). htmlContent là khung phiếu chứa vùng kết quả.
type StaffCreateImagingResultSheetTemplateRequest struct {
	TeleradPartnerUuid uuid.UUID `json:"teleradPartnerUuid" validate:"required"`
	HtmlContent        string    `json:"htmlContent" validate:"required"`
}
