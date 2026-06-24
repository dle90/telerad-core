package objectMappers

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	databaseQueryModels "telerad-core-module/internals/models/database-query_models"
	teleradReadingOrderControllerResponses "telerad-core-module/internals/responses/telerad-reading-order-controller_responses"
	"telerad-core-module/internals/types"
)

// ToReadingOrderPartnerGroupSlice gom đối tác thành cây "loại chụp → đối tác" cho
// màn "Đọc ca". Chỉ gom các loại chụp có trong allowedModalities (ADMIN truyền vào
// toàn bộ MODALITIES, user thường truyền các loại chụp được phân). Nhóm hiển thị
// theo thứ tự chuẩn trong constants.MODALITIES; 1 đối tác xuất hiện ở mọi loại chụp
// nó cung cấp; bỏ qua nhóm loại chụp không có đối tác nào.
func ToReadingOrderPartnerGroupSlice(
	partners []entities.TeleradPartnerEntity,
	allowedModalities []string,
) []teleradReadingOrderControllerResponses.ReadingOrderPartnerGroupResponse {
	allowed := make(map[string]bool, len(allowedModalities))
	for _, m := range allowedModalities {
		allowed[m] = true
	}

	groups := make([]teleradReadingOrderControllerResponses.ReadingOrderPartnerGroupResponse, 0, len(constants.MODALITIES))
	for _, modality := range constants.MODALITIES {
		if !allowed[modality] {
			continue
		}

		items := make([]teleradReadingOrderControllerResponses.ReadingOrderPartnerItem, 0)
		for _, p := range partners {
			if !slices.Contains(p.Modalities, modality) {
				continue
			}
			items = append(items, teleradReadingOrderControllerResponses.ReadingOrderPartnerItem{
				Uuid: p.Uuid,
				Code: p.Code,
				Name: p.Name,
			})
		}

		if len(items) == 0 {
			continue
		}

		groups = append(groups, teleradReadingOrderControllerResponses.ReadingOrderPartnerGroupResponse{
			Modality: modality,
			Partners: items,
		})
	}

	return groups
}

func ToPartnerCreateReadingOrderResponse(readingOrder entities.TeleradReadingOrderEntity) teleradReadingOrderControllerResponses.PartnerCreateReadingOrderResponse {
	return teleradReadingOrderControllerResponses.PartnerCreateReadingOrderResponse{
		Uuid: readingOrder.Uuid,
	}
}

// ToStaffGetReadingOrderDetailResponse map entity ca đọc + tên đối tác + tên bác sĩ
// đọc (đã resolve ở service) sang response chi tiết cho tab chi tiết.
func ToStaffGetReadingOrderDetailResponse(
	ro entities.TeleradReadingOrderEntity,
	partnerName string,
	assignedToName *string,
	assignedToMe bool,
) teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse {
	return teleradReadingOrderControllerResponses.StaffGetReadingOrderDetailResponse{
		Uuid:               ro.Uuid,
		TeleradPartnerUuid: ro.TeleradPartnerUuid,
		PartnerName:        partnerName,
		OrderCode:          ro.OrderCode,
		OrderItemCode:      ro.OrderItemCode,
		StudyInstanceUid:   ro.StudyInstanceUid,
		PatientCode:        ro.PatientCode,
		FullName:           ro.FullName,
		DateOfBirth:        ro.DateOfBirth,
		Gender:             ro.Gender,
		YearsOld:           ro.YearsOld,
		MonthsOld:          ro.MonthsOld,
		DaysOld:            ro.DaysOld,
		Phone:              ro.Phone,
		Email:              ro.Email,
		FullAddress:        ro.FullAddress,
		ServiceName:        ro.ServiceName,
		Modality:           ro.Modality,
		ModalityName:       ro.ModalityName,
		Note:               ro.Note,
		PerformEndedAt:     ro.PerformEndedAt,
		ClinicalDiagnosis:  ro.ClinicalDiagnosis,
		Icd:                ro.Icd,
		BodyParts:          ro.BodyParts,
		AssignedTo:         ro.AssignedTo,
		AssignedToName:     assignedToName,
		AssignedToMe:       assignedToMe,
		ReadCompletedAt:    ro.ReadCompletedAt,
		Status:             ro.Status,
		ResultReturned:     ro.ResultReturned,
		ResultInHtml:       ro.ResultInHtml,
	}
}

// ToPublicGetReadingOrderResultSheetResponse lắp dữ liệu in cho mẫu phiếu kết quả: key của
// data = ĐÚNG tên token trên mẫu, đã format hiển thị sẵn (giới tính, ngày, năm sinh, cỡ chữ)
// để frontend chỉ cần fillTokens. readBy (tên bác sĩ đọc) đã resolve ở service.
func ToPublicGetReadingOrderResultSheetResponse(
	ro entities.TeleradReadingOrderEntity,
	htmlContent string,
	resultFontSize int16,
	resultLineSpacing float64,
	readBy string,
) teleradReadingOrderControllerResponses.PublicGetReadingOrderResultSheetResponse {
	// Ngày trên phiếu: nếu ca chưa đọc xong (ReadCompletedAt null) -> lấy thời gian hiện tại (ngày in).
	readCompletedAt := ro.ReadCompletedAt
	if readCompletedAt == nil || readCompletedAt.IsZero() {
		now := time.Now()
		readCompletedAt = &now
	}

	return teleradReadingOrderControllerResponses.PublicGetReadingOrderResultSheetResponse{
		HtmlContent:       htmlContent,
		ResultFontSize:    resultFontSize,
		ResultLineSpacing: resultLineSpacing,
		Data: teleradReadingOrderControllerResponses.ReadingOrderPrintData{
			PatientName:       ro.FullName,
			PatientBirthYear:  printBirthYear(ro.DateOfBirth),
			PatientGender:     printGenderLabel(ro.Gender),
			IndicationPlace:   "",
			ServiceName:       ro.ServiceName,
			ClinicalDiagnosis: derefString(ro.ClinicalDiagnosis),
			ResultContent:     derefString(ro.ResultInHtml),
			ReadCompletedAt:   printVietnameseDate(readCompletedAt),
			ReadBy:            readBy,
			LogoUrl:           "",
		},
	}
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// printGenderLabel chuẩn hoá giới tính sang "Nam"/"Nữ" để in.
func printGenderLabel(g *string) string {
	if g == nil {
		return ""
	}
	switch strings.ToUpper(strings.TrimSpace(*g)) {
	case "MALE", "M", "NAM":
		return "Nam"
	case "FEMALE", "F", "NU", "NỮ":
		return "Nữ"
	default:
		return *g
	}
}

func printBirthYear(d *types.Date) string {
	if d == nil || d.IsZero() {
		return ""
	}
	return strconv.Itoa(d.Time().Year())
}

// printVietnameseDate -> "24 tháng 06 năm 2026" (rỗng nếu chưa có ngày).
func printVietnameseDate(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d tháng %02d năm %d", t.Day(), int(t.Month()), t.Year())
}

func ToStaffGetListReadingOrderSlice(
	rows []databaseQueryModels.ReadingOrderListRow,
) []teleradReadingOrderControllerResponses.StaffGetListReadingOrderResponse {
	result := make([]teleradReadingOrderControllerResponses.StaffGetListReadingOrderResponse, 0, len(rows))

	for _, row := range rows {
		result = append(result, teleradReadingOrderControllerResponses.StaffGetListReadingOrderResponse{
			Uuid:               row.Uuid,
			TeleradPartnerUuid: row.TeleradPartnerUuid,
			PartnerName:        row.PartnerName,
			OrderCode:          row.OrderCode,
			OrderItemCode:      row.OrderItemCode,
			StudyInstanceUid:   row.StudyInstanceUid,
			PatientCode:        row.PatientCode,
			FullName:           row.FullName,
			Gender:             row.Gender,
			YearsOld:           row.YearsOld,
			Phone:              row.Phone,
			ServiceName:        row.ServiceName,
			Modality:           row.Modality,
			ModalityName:       row.ModalityName,
			PerformEndedAt:     row.PerformEndedAt,
			ReadCompletedAt:    row.ReadCompletedAt,
			AssignedTo:         row.AssignedTo,
			AssignedToName:     row.AssignedToName,
			Status:             row.Status,
			ResultReturned:     row.ResultReturned,
		})
	}

	return result
}
