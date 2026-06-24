package errors

// ============================ TELERAD CORE ============================

// Nhóm lỗi TELERAD E001: chung / xác thực
const TELERAD_E001_001 string = "Mật khẩu hiện tại không đúng"
const TELERAD_E001_002 string = "Tài khoản của bạn đã bị khóa"
const TELERAD_E001_003 string = "Tài khoản hoặc mật khẩu không đúng"

// Nhóm lỗi TELERAD E101: đối tác telerad (telerad_partner)
const TELERAD_E101_001 string = "Đối tác telerad không tồn tại"
const TELERAD_E101_002 string = "Mã đối tác đã tồn tại"
const TELERAD_E101_003 string = "Tên đăng nhập (phía telerad) đã tồn tại"

// Nhóm lỗi TELERAD E102: tài khoản nhân viên (staff_account)
const TELERAD_E102_001 string = "Nhân viên không tồn tại"
const TELERAD_E102_002 string = "Mã nhân viên đã tồn tại"
const TELERAD_E102_003 string = "Tên đăng nhập đã tồn tại"
const TELERAD_E102_004 string = "Nhân viên đã có tài khoản đăng nhập"
const TELERAD_E102_005 string = "Nhân viên chưa có tài khoản đăng nhập"
const TELERAD_E102_006 string = "Không thể thao tác trên tài khoản quản trị"

// Nhóm lỗi TELERAD E103: ca đọc (telerad_reading_order)
const TELERAD_E103_001 string = "Ca đọc không tồn tại"
const TELERAD_E103_002 string = "Bạn không có quyền đọc ca này"
const TELERAD_E103_003 string = "Ca đọc không ở trạng thái chưa đọc"
const TELERAD_E103_004 string = "Ca đọc không ở trạng thái đang đọc của bạn"
const TELERAD_E103_005 string = "Bạn đang đọc ca %s của bệnh nhân %s" // format: order_item_code, full_name
const TELERAD_E103_006 string = "Chưa có nội dung kết quả để duyệt"

// Nhóm lỗi TELERAD E104: mẫu kết quả (imaging_result_template)
const TELERAD_E104_001 string = "Mẫu kết quả không tồn tại"
const TELERAD_E104_002 string = "Loại chụp không hợp lệ"
const TELERAD_E104_003 string = "Bộ phận chụp không hợp lệ"

// Nhóm lỗi TELERAD E105: phiếu kết quả (imaging_result_sheet_template)
const TELERAD_E105_001 string = "Phiếu kết quả không tồn tại"
const TELERAD_E105_002 string = "Cơ sở y tế này đã có phiếu kết quả"
