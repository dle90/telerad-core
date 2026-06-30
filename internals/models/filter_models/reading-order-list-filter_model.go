package filterModels

import (
	"time"

	"github.com/google/uuid"
)

// ReadingOrderListFilter gom tham số lọc danh sách ca đọc màn "Đọc ca".
// ScopePartnerUuids / ScopeModalities là phạm vi QUYỀN của user: nil = không giới
// hạn (ADMIN); khác nil = chỉ lấy ca thuộc các partner / loại chụp này.
type ReadingOrderListFilter struct {
	IsAdmin          bool        // có phải admin không (quyền xem tất cả partner + modality)
	PartnerUuids     []uuid.UUID // quyền: partner được đọc (nil = admin, không giới hạn)
	Modalities       []string    // quyền: loại chụp được đọc (nil = admin, không giới hạn)
	PerformEndedFrom *time.Time  // lọc theo ngày chụp (perform_ended_at) — từ
	PerformEndedTo   *time.Time  // lọc theo ngày chụp — đến
	PatientName      string      // tên bệnh nhân (ILIKE)
	PatientCode      string      // mã bệnh nhân (ILIKE)
	Phone            string      // số điện thoại (ILIKE)
	Status           string      // tình trạng ca (status) — "" = tất cả
	AssignedTo       *uuid.UUID  // lọc theo bác sĩ đang nhận (assigned_to) — nil = không lọc
	ResultReturned   *bool       // đã trả kết quả chưa — nil = tất cả
}
