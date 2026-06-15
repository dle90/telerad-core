package utils

import "time"

// GMT7 là múi giờ Việt Nam (UTC+7, không có DST từ 1975). Dùng làm baseline
// cho các tính toán dựa trên lịch VN — reset sequence theo năm, group thống
// kê theo ngày/tháng — để kết quả không phụ thuộc múi giờ server (container
// thường chạy UTC).
//
// Dùng FixedZone thay vì time.LoadLocation("Asia/Ho_Chi_Minh") để không phụ
// thuộc tzdata trên image runtime (distroless / scratch không có).
var GMT7 = time.FixedZone("GMT+7", 7*60*60)

// NowGMT7 trả về thời gian hiện tại quy về múi giờ GMT+7. Dùng khi cần lấy
// Year/Month/Day/Hour... theo lịch Việt Nam.
func NowGMT7() time.Time {
	return time.Now().In(GMT7)
}

func NowGMT7_year() int {
	return NowGMT7().Year()
}
