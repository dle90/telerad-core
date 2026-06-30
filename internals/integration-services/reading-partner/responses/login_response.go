package partnerResponses

// LoginResponse — token đối tác (his-core) trả về sau khi telerad đăng nhập
// integration. Khớp his-core GetTokenResponse (access_token/token_type/expires_in).
type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"` // TTL của token tính bằng giây
}
