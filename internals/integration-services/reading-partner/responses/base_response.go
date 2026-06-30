package partnerResponses

// BaseResponse — vỏ phản hồi chung của đối tác (his-core): { code, message, result }.
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
