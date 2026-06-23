package imagingResultTemplateControllerResponses

// StaffGetImagingResultTemplateFormOptionsResponse — dữ liệu cho form mẫu kết quả:
// danh sách loại chụp hỗ trợ + danh sách bộ phận chụp (PACS_BODY_PART) để tick chọn.
type StaffGetImagingResultTemplateFormOptionsResponse struct {
	Modalities []string                      `json:"modalities"`
	BodyParts  []ImagingResultTemplateOption `json:"bodyParts"`
}

type ImagingResultTemplateOption struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
