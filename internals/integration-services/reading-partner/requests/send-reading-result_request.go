package partnerRequests

// SendReadingResultRequest — payload telerad trả kết quả 1 ca đã DUYỆT về đối tác
// (his-core). Đối tác đối chiếu lại chỉ định bằng orderItemId (= uuid chỉ định CĐHA
// phía his-core, đúng giá trị his-core đã gửi sang khi tạo ca) + teleradReadingOrderUuid.
type SendReadingResultRequest struct {
	// định danh ca phía telerad
	TeleradReadingOrderUuid string `json:"teleradReadingOrderUuid"`

	// định danh ca từ phía đối tác (echo lại đúng giá trị đối tác đã gửi)
	OrderId          string `json:"orderId"`
	OrderCode        string `json:"orderCode"`
	OrderItemId      string `json:"orderItemId"`
	OrderItemCode    string `json:"orderItemCode"`
	StudyInstanceUid string `json:"studyInstanceUid"`

	// kết quả
	Status         string  `json:"status"`         // APPROVED
	ResultInHtml   *string `json:"resultInHtml"`   // nội dung kết quả (HTML)
	ResultInText   *string `json:"resultInText"`   // nội dung kết quả (text thuần)
	ResultedAt     string  `json:"resultedAt"`     // RFC3339 = approved_at
	ApprovedByName string  `json:"approvedByName"` // tên bác sĩ duyệt
}
