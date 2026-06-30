package teleradReadingOrderControllerResponses

// StaffEndReadingAndApproveResponse — kết quả nút "Kết thúc & Duyệt": chi tiết ca sau khi
// duyệt + cờ cho biết việc TRẢ KẾT QUẢ về đối tác có thất bại không.
//
// resultReturnFailed = false khi: đối tác KHÔNG bật callback (không cần trả), HOẶC trả kết
// quả thành công. = true khi: đã duyệt OK nhưng đẩy kết quả sang đối tác thất bại → FE báo
// "duyệt thành công nhưng trả kết quả thất bại" (ca vẫn APPROVED, có thể bấm "Trả KQ" gửi lại).
type StaffEndReadingAndApproveResponse struct {
	ReadingOrder       *StaffGetReadingOrderDetailResponse `json:"readingOrder"`
	ResultReturnFailed bool                                `json:"resultReturnFailed"`
}
