package baseServices

import (
	"context"
	"time"

	"telerad-core-module/internals/entities"
	"telerad-core-module/internals/repositories"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindOneImagingResultSheetTemplateByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.ImagingResultSheetTemplateEntity, error) {
	return repositories.FindOneByUuid[entities.ImagingResultSheetTemplateEntity](ctx, tx, id)
}

func InitNewImagingResultSheetTemplate(teleradPartnerUuid uuid.UUID, htmlContent string) entities.ImagingResultSheetTemplateEntity {
	return entities.ImagingResultSheetTemplateEntity{
		TeleradPartnerUuid: teleradPartnerUuid,
		HtmlContent:        htmlContent,
		IsActive:           true,
	}
}

func CreateNewImagingResultSheetTemplate(ctx context.Context, tx bun.IDB, creatorUuid uuid.UUID, newRecord *entities.ImagingResultSheetTemplateEntity) error {
	newRecord.CreatedAt = time.Now()
	newRecord.CreatedBy = creatorUuid

	return repositories.InsertOne(ctx, tx, newRecord)
}

func UpdateWholeImagingResultSheetTemplateRecord(ctx context.Context, tx bun.IDB, updaterUuid uuid.UUID, template *entities.ImagingResultSheetTemplateEntity) error {
	now := time.Now()
	template.UpdatedAt = &now
	template.UpdatedBy = &updaterUuid

	return repositories.UpdateWholeRecord(ctx, tx, template)
}
