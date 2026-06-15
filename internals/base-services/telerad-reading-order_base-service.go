package baseServices

import (
	"context"
	"time"

	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	"telerad-core-module/internals/repositories"
	teleradReadingOrderControllerRequests "telerad-core-module/internals/requests/telerad-reading-order-controller_requests"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindOneTeleradReadingOrderByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.TeleradReadingOrderEntity, error) {
	return repositories.FindOneByUuid[entities.TeleradReadingOrderEntity](ctx, tx, id)
}

// InitNewTeleradReadingOrder dựng ca đọc mới từ payload đối tác. status = PENDING,
// các trường phân công (assigned_*, read_completed_at) để trống.
func InitNewTeleradReadingOrder(teleradPartnerUuid uuid.UUID, request teleradReadingOrderControllerRequests.PartnerCreateReadingOrderRequest) entities.TeleradReadingOrderEntity {
	// Tuổi tính tại thời điểm thực hiện chụp (perform_ended_at).
	years, months, days := calculateAgeBreakdown(request.DateOfBirth.Time(), request.PerformEndedAt)

	return entities.TeleradReadingOrderEntity{
		TeleradPartnerUuid: teleradPartnerUuid,
		OrderId:            request.OrderId,
		OrderCode:          request.OrderCode,
		OrderItemId:        request.OrderItemId,
		OrderItemCode:      request.OrderItemCode,
		StudyInstanceUid:   request.StudyInstanceUid,
		PatientCode:        request.PatientCode,
		FullName:           request.FullName,
		DateOfBirth:        &request.DateOfBirth,
		Gender:             request.Gender,
		Phone:              request.Phone,
		Email:              request.Email,
		FullAddress:        request.FullAddress,
		YearsOld:           &years,
		MonthsOld:          &months,
		DaysOld:            &days,
		ServiceId:          request.ServiceId,
		ServiceCode:        request.ServiceCode,
		ServiceName:        request.ServiceName,
		Modality:           &request.Modality,
		ModalityAeTitle:    request.ModalityAeTitle,
		ModalityCode:       request.ModalityCode,
		ModalityName:       request.ModalityName,
		Note:               request.Note,
		PerformEndedAt:     request.PerformEndedAt,
		ClinicalDiagnosis:  request.ClinicalDiagnosis,
		Icd:                request.Icd,
		Status:             entities.TeleradReadingOrderStatusPending,
	}
}

func CreateNewTeleradReadingOrder(ctx context.Context, tx bun.IDB, creatorUuid uuid.UUID, newRecord *entities.TeleradReadingOrderEntity) error {
	newRecord.CreatedAt = time.Now()
	newRecord.CreatedBy = creatorUuid

	return repositories.InsertOne(ctx, tx, newRecord)
}

// OverwriteTeleradReadingOrderInfo ghi đè thông tin ca đọc từ payload đối tác lên bản ghi
// đã tồn tại (đối tác gửi lại cùng order item). GIỮ NGUYÊN trạng thái + thông tin phân công
// (status, assigned_at/assigned_to, read_completed_at) — chỉ làm mới phần thông tin
// order / bệnh nhân / dịch vụ / lâm sàng. Tuổi tính lại theo dob + perform_ended_at mới.
func OverwriteTeleradReadingOrderInfo(record *entities.TeleradReadingOrderEntity, request teleradReadingOrderControllerRequests.PartnerCreateReadingOrderRequest) {
	years, months, days := calculateAgeBreakdown(request.DateOfBirth.Time(), request.PerformEndedAt)

	record.OrderId = request.OrderId
	record.OrderCode = request.OrderCode
	record.OrderItemCode = request.OrderItemCode
	record.StudyInstanceUid = request.StudyInstanceUid
	record.PatientCode = request.PatientCode
	record.FullName = request.FullName
	record.DateOfBirth = &request.DateOfBirth
	record.Gender = request.Gender
	record.Phone = request.Phone
	record.Email = request.Email
	record.FullAddress = request.FullAddress
	record.YearsOld = &years
	record.MonthsOld = &months
	record.DaysOld = &days
	record.ServiceId = request.ServiceId
	record.ServiceCode = request.ServiceCode
	record.ServiceName = request.ServiceName
	record.Modality = &request.Modality
	record.ModalityAeTitle = request.ModalityAeTitle
	record.ModalityCode = request.ModalityCode
	record.ModalityName = request.ModalityName
	record.Note = request.Note
	record.PerformEndedAt = request.PerformEndedAt
	record.ClinicalDiagnosis = request.ClinicalDiagnosis
	record.Icd = request.Icd
}

func UpdateWholeTeleradReadingOrderRecord(ctx context.Context, tx bun.IDB, updaterUuid uuid.UUID, record *entities.TeleradReadingOrderEntity) error {
	now := time.Now()
	record.UpdatedAt = &now
	record.UpdatedBy = &updaterUuid

	return repositories.UpdateWholeRecord(ctx, tx, record)
}

// calculateAgeBreakdown trả về tuổi giữa dob và ref ở 3 mức: số năm tròn, tổng số tháng,
// tổng số ngày. Cả hai mốc đều chuẩn hoá về GMT+7 để lấy đúng ngày theo lịch Việt Nam.
// Ví dụ: dob=2020-03-10, ref=2026-05-12 → years=6, months=74, days=2254.
func calculateAgeBreakdown(dob time.Time, ref time.Time) (years int16, months int16, days int) {
	dob = dob.In(constants.GMT7_TIMEZONE)
	ref = ref.In(constants.GMT7_TIMEZONE)

	y := ref.Year() - dob.Year()
	m := int(ref.Month()) - int(dob.Month())
	if ref.Day() < dob.Day() {
		m--
	}
	if m < 0 {
		y--
		m += 12
	}
	totalMonths := y*12 + m

	dobDate := time.Date(dob.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, dob.Location())
	refDate := time.Date(ref.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, ref.Location())
	totalDays := int(refDate.Sub(dobDate) / (24 * time.Hour))

	return int16(y), int16(totalMonths), totalDays
}
