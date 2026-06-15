package objectMappers

import (
	"telerad-core-module/internals/entities"
	staffAccountControllerResponses "telerad-core-module/internals/responses/staff-account-controller_responses"
)

func ToStaffGetListStaffAccountSlice(staffs []entities.StaffAccountEntity) []staffAccountControllerResponses.StaffGetListStaffAccountResponse {
	result := make([]staffAccountControllerResponses.StaffGetListStaffAccountResponse, 0, len(staffs))

	for _, staff := range staffs {
		result = append(result, staffAccountControllerResponses.StaffGetListStaffAccountResponse{
			Uuid:       staff.Uuid,
			Code:       staff.Code,
			FullName:   staff.FullName,
			Gender:     staff.Gender,
			Phone:      staff.Phone,
			Email:      staff.Email,
			Username:   staff.Username,
			IsActive:   staff.IsActive,
			Modalities: emptyIfNil(staff.Modalities),
			Roles:      emptyIfNil(staff.Roles),
			CreatedAt:  staff.CreatedAt,
			HasAccount: staff.Username != nil,
		})
	}

	return result
}

func ToStaffGetAStaffAccountResponse(staff entities.StaffAccountEntity) staffAccountControllerResponses.StaffGetAStaffAccountResponse {
	return staffAccountControllerResponses.StaffGetAStaffAccountResponse{
		Uuid:                  staff.Uuid,
		CreatedAt:             staff.CreatedAt,
		CreatedBy:             staff.CreatedBy,
		UpdatedAt:             staff.UpdatedAt,
		UpdatedBy:             staff.UpdatedBy,
		Code:                  staff.Code,
		FullName:              staff.FullName,
		DateOfBirth:           staff.DateOfBirth,
		Gender:                staff.Gender,
		CitizenIdentityNumber: staff.CitizenIdentityNumber,
		Phone:                 staff.Phone,
		Email:                 staff.Email,
		FullAddress:           staff.FullAddress,
		IsActive:              staff.IsActive,
		Username:              staff.Username,
		Modalities:            emptyIfNil(staff.Modalities),
		Roles:                 emptyIfNil(staff.Roles),
		TeleradPartnerUuids:   emptyIfNil(staff.TeleradPartnerUuids),
		HasAccount:            staff.Username != nil,
	}
}

func ToUserGetMeResponse(staff entities.StaffAccountEntity) staffAccountControllerResponses.UserGetMeResponse {
	return staffAccountControllerResponses.UserGetMeResponse{
		Uuid:                  staff.Uuid,
		Code:                  staff.Code,
		FullName:              staff.FullName,
		DateOfBirth:           staff.DateOfBirth,
		Gender:                staff.Gender,
		CitizenIdentityNumber: staff.CitizenIdentityNumber,
		Phone:                 staff.Phone,
		Email:                 staff.Email,
		FullAddress:           staff.FullAddress,
		Username:              staff.Username,
		IsActive:              staff.IsActive,
		Modalities:            emptyIfNil(staff.Modalities),
		Roles:                 emptyIfNil(staff.Roles),
		TeleradPartnerUuids:   emptyIfNil(staff.TeleradPartnerUuids),
	}
}
