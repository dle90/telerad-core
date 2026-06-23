package objectMappers

import (
	"slices"

	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	databaseQueryModels "telerad-core-module/internals/models/database-query_models"
	teleradReadingOrderControllerResponses "telerad-core-module/internals/responses/telerad-reading-order-controller_responses"
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
		})
	}

	return result
}
