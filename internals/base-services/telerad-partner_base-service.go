package baseServices

import (
	"context"
	"strings"
	"time"

	"telerad-core-module/internals/entities"
	"telerad-core-module/internals/repositories"
	teleradPartnerControllerRequests "telerad-core-module/internals/requests/telerad-partner-controller_requests"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindOneTeleradPartnerByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.TeleradPartnerEntity, error) {
	return repositories.FindOneByUuid[entities.TeleradPartnerEntity](ctx, tx, id)
}

// GetTeleradPartnerByUsername — uppercase username trước khi query (username lưu uppercase).
func GetTeleradPartnerByUsername(ctx context.Context, tx bun.IDB, username string) (*entities.TeleradPartnerEntity, error) {
	return repositories.FindOneTeleradPartnerByUsername(ctx, tx, strings.ToUpper(username), nil)
}

// InitNewTeleradPartner dựng entity mới từ request. username + passwordHash đã được
// service chuẩn hoá (uppercase username, bcrypt hash). Status mặc định ACTIVE.
func InitNewTeleradPartner(request teleradPartnerControllerRequests.StaffCreateTeleradPartnerRequest, username, passwordHash string) entities.TeleradPartnerEntity {
	return entities.TeleradPartnerEntity{
		Code:            request.Code,
		Name:            request.Name,
		IsActive:        true,
		Contact:         request.Contact,
		Callback:        request.Callback,
		CallbackUrl:     request.CallbackUrl,
		Username:        username,
		PasswordHash:    passwordHash,
		PartnerUsername: request.PartnerUsername,
		PartnerPassword: request.PartnerPassword,
		Modalities:      request.Modalities,
	}
}

func CreateNewTeleradPartner(ctx context.Context, tx bun.IDB, creatorUuid uuid.UUID, newRecord *entities.TeleradPartnerEntity) error {
	newRecord.CreatedAt = time.Now()
	newRecord.CreatedBy = creatorUuid

	return repositories.InsertOne(ctx, tx, newRecord)
}

func UpdateWholeTeleradPartnerRecord(ctx context.Context, tx bun.IDB, updaterUuid uuid.UUID, partner *entities.TeleradPartnerEntity) error {
	now := time.Now()
	partner.UpdatedAt = &now
	partner.UpdatedBy = &updaterUuid

	return repositories.UpdateWholeRecord(ctx, tx, partner)
}
