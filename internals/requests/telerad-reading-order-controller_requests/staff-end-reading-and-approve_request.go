package teleradReadingOrderControllerRequests

// StaffEndReadingAndApproveRequest — body nút "Kết thúc & Duyệt": lưu nội dung kết quả
// (html) rồi duyệt ca. resultInHtml bắt buộc không rỗng (BE kiểm tra lại dù FE đã chặn).
type StaffEndReadingAndApproveRequest struct {
	ResultInHtml string `json:"resultInHtml"`
}
