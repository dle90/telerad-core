package objectMappers

import (
	"telerad-core-module/internals/entities"
	teleradPartnerControllerResponses "telerad-core-module/internals/responses/telerad-partner-controller_responses"
)

func ToStaffCreateTeleradPartnerResponse(partner entities.TeleradPartnerEntity) teleradPartnerControllerResponses.StaffCreateTeleradPartnerResponse {
	return teleradPartnerControllerResponses.StaffCreateTeleradPartnerResponse{
		Uuid:     partner.Uuid,
		Code:     partner.Code,
		Name:     partner.Name,
		IsActive: partner.IsActive,
		Username: partner.Username,
	}
}

func ToStaffGetListTeleradPartnerSlice(partners []entities.TeleradPartnerEntity) []teleradPartnerControllerResponses.StaffGetListTeleradPartnerResponse {
	result := make([]teleradPartnerControllerResponses.StaffGetListTeleradPartnerResponse, 0, len(partners))

	for _, partner := range partners {
		result = append(result, teleradPartnerControllerResponses.StaffGetListTeleradPartnerResponse{
			Uuid:       partner.Uuid,
			Code:       partner.Code,
			Name:       partner.Name,
			IsActive:   partner.IsActive,
			Username:   partner.Username,
			Contact:    partner.Contact,
			Callback:   partner.Callback,
			Modalities: emptyIfNil(partner.Modalities),
			CreatedAt:  partner.CreatedAt,
		})
	}

	return result
}

func ToStaffGetAllTeleradPartnerSlice(partners []entities.TeleradPartnerEntity) []teleradPartnerControllerResponses.StaffGetAllTeleradPartnerResponse {
	result := make([]teleradPartnerControllerResponses.StaffGetAllTeleradPartnerResponse, 0, len(partners))

	for _, partner := range partners {
		result = append(result, teleradPartnerControllerResponses.StaffGetAllTeleradPartnerResponse{
			Uuid:       partner.Uuid,
			Code:       partner.Code,
			Name:       partner.Name,
			IsActive:   partner.IsActive,
			Modalities: emptyIfNil(partner.Modalities),
		})
	}

	return result
}

func ToStaffGetATeleradPartnerResponse(partner entities.TeleradPartnerEntity) teleradPartnerControllerResponses.StaffGetATeleradPartnerResponse {
	return teleradPartnerControllerResponses.StaffGetATeleradPartnerResponse{
		Uuid:            partner.Uuid,
		CreatedAt:       partner.CreatedAt,
		CreatedBy:       partner.CreatedBy,
		UpdatedAt:       partner.UpdatedAt,
		UpdatedBy:       partner.UpdatedBy,
		Code:            partner.Code,
		Name:            partner.Name,
		IsActive:        partner.IsActive,
		Contact:         partner.Contact,
		Username:        partner.Username,
		Callback:        partner.Callback,
		CallbackUrl:     partner.CallbackUrl,
		PartnerUsername: partner.PartnerUsername,
		Modalities:      emptyIfNil(partner.Modalities),
	}
}

func ToStaffTeleradPartnerPartnerConfigResponse(partner entities.TeleradPartnerEntity) teleradPartnerControllerResponses.StaffTeleradPartnerPartnerConfigResponse {
	return teleradPartnerControllerResponses.StaffTeleradPartnerPartnerConfigResponse{
		Uuid:            partner.Uuid,
		Callback:        partner.Callback,
		CallbackUrl:     partner.CallbackUrl,
		PartnerUsername: partner.PartnerUsername,
		PartnerPassword: partner.PartnerPassword,
	}
}
