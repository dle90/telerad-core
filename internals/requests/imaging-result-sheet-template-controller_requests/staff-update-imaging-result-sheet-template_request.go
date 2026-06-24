package imagingResultSheetTemplateControllerRequests

// StaffUpdateImagingResultSheetTemplateRequest — sửa nội dung mẫu phiếu kết quả.
// Không đổi CSYT (telerad_partner cố định theo bản ghi); status có endpoint riêng.
type StaffUpdateImagingResultSheetTemplateRequest struct {
	HtmlContent string `json:"htmlContent" validate:"required"`
	// Cỡ chữ + giãn dòng áp cho vùng kết quả khi IN phiếu.
	ResultFontSize    int16   `json:"resultFontSize" validate:"required,min=1"`
	ResultLineSpacing float64 `json:"resultLineSpacing" validate:"required,gt=0"`
}
