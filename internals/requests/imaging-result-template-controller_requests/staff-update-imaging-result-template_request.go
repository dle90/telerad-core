package imagingResultTemplateControllerRequests

// StaffUpdateImagingResultTemplateRequest — sửa mẫu nội dung kết quả. Không đụng status
// (có endpoint activate/deactivate riêng).
type StaffUpdateImagingResultTemplateRequest struct {
	Modality     string   `json:"modality" validate:"required,oneof=CT MR US CR MG"`
	Name         string   `json:"name" validate:"required"`
	BodyParts    []string `json:"bodyParts" validate:"omitempty,dive,required"`
	HtmlContent  string   `json:"htmlContent" validate:"required"`
	FontSize     int      `json:"fontSize" validate:"required,min=1"`
	LineSpacing  float64  `json:"lineSpacing" validate:"required,gt=0"`
	DisplayOrder *int16   `json:"displayOrder"`
}
