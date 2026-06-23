package teleradReadingOrderControllerResponses

import "github.com/google/uuid"

// ReadingOrderPartnerGroupResponse — 1 nhóm "loại chụp" ở cây bên trái màn "Đọc
// ca", chứa danh sách đối tác (đã lọc theo quyền của user) cung cấp loại chụp đó.
type ReadingOrderPartnerGroupResponse struct {
	Modality string                    `json:"modality"`
	Partners []ReadingOrderPartnerItem `json:"partners"`
}

type ReadingOrderPartnerItem struct {
	Uuid uuid.UUID `json:"uuid"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}
