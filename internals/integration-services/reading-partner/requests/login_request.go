package partnerRequests

// LoginRequest — payload telerad đăng nhập vào đối tác (his-core) để lấy token
// trước khi trả kết quả. Username/Password = PartnerUsername/PartnerPassword đã
// cấu hình trên telerad_partner.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
