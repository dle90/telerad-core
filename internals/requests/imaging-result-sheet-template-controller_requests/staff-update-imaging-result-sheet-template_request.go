package imagingResultSheetTemplateControllerRequests

// StaffUpdateImagingResultSheetTemplateRequest — sửa nội dung mẫu phiếu kết quả.
// Không đổi CSYT (telerad_partner cố định theo bản ghi); status có endpoint riêng.
type StaffUpdateImagingResultSheetTemplateRequest struct {
	HtmlContent string `json:"htmlContent" validate:"required"`
}
