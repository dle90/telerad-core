package imagingResultTemplateControllerRequests

// StaffCreateImagingResultTemplateRequest — tạo mẫu nội dung kết quả CĐHA.
// modality giới hạn theo loại chụp hệ thống hỗ trợ (CT MR US CR MG); bodyParts là
// mã PACS_BODY_PART (validate sâu ở service).
type StaffCreateImagingResultTemplateRequest struct {
	Modality     string   `json:"modality" validate:"required,oneof=CT MR US CR MG"`
	Name         string   `json:"name" validate:"required"`
	BodyParts    []string `json:"bodyParts" validate:"omitempty,dive,required"`
	HtmlContent  string   `json:"htmlContent" validate:"required"`
	FontSize     int      `json:"fontSize" validate:"required,min=1"`
	LineSpacing  float64  `json:"lineSpacing" validate:"required,gt=0"`
	DisplayOrder *int16   `json:"displayOrder"`
}
