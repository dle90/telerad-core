package teleradReadingOrderControllerResponses

// StaffGetReadingOrderResultSheetResponse — mẫu phiếu kết quả (của CSYT ca đọc) để dựng bản in.
//   htmlContent: HTML mẫu phiếu (chứa token {{...}}).
//   resultFontSize/resultLineSpacing: CẤU HÌNH của phiếu (cột result_font_size/line_spacing),
//     áp cho vùng kết quả khi in — để top-level (typed) vì là config, không phải token-data ca đọc.
//   data: giá trị token NỘI DUNG ca đọc, KEY = ĐÚNG TÊN TOKEN -> frontend fillTokens, không map.
type StaffGetReadingOrderResultSheetResponse struct {
	HtmlContent       string                `json:"htmlContent"`
	ResultFontSize    int16                 `json:"resultFontSize"`
	ResultLineSpacing float64               `json:"resultLineSpacing"`
	Data              ReadingOrderPrintData `json:"data"`
}

// ReadingOrderPrintData — giá trị token NỘI DUNG của mẫu phiếu kết quả. Backend lo format hiển
// thị (giới tính Nam/Nữ, ngày tiếng Việt, năm sinh). Mọi giá trị là chuỗi để thay token trực
// tiếp. Thêm token nội dung mới trên mẫu -> thêm field tương ứng ở đây.
type ReadingOrderPrintData struct {
	PatientName       string `json:"patientName"`
	PatientBirthYear  string `json:"patientBirthYear"`
	PatientGender     string `json:"patientGender"`
	IndicationPlace   string `json:"indicationPlace"`
	ServiceName       string `json:"serviceName"`
	ClinicalDiagnosis string `json:"clinicalDiagnosis"`
	ResultContent     string `json:"resultContent"`
	ReadCompletedAt   string `json:"readCompletedAt"`
	ReadBy            string `json:"readBy"`
	LogoUrl           string `json:"logoUrl"`
}
