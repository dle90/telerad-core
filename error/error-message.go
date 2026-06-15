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
