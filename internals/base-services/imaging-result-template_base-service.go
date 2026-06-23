package baseServices

import (
	"context"
	"time"

	"telerad-core-module/internals/entities"
	"telerad-core-module/internals/repositories"
	imagingResultTemplateControllerRequests "telerad-core-module/internals/requests/imaging-result-template-controller_requests"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindOneImagingResultTemplateByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.ImagingResultTemplateEntity, error) {
	return repositories.FindOneByUuid[entities.ImagingResultTemplateEntity](ctx, tx, id)
}

func InitNewImagingResultTemplate(request imagingResultTemplateControllerRequests.StaffCreateImagingResultTemplateRequest) entities.ImagingResultTemplateEntity {
	return entities.ImagingResultTemplateEntity{
		Modality:     request.Modality,
		Name:         request.Name,
		BodyParts:    request.BodyParts,
		HtmlContent:  request.HtmlContent,
		FontSize:     request.FontSize,
		LineSpacing:  request.LineSpacing,
		DisplayOrder: request.DisplayOrder,
		IsActive:     true,
	}
}

func CreateNewImagingResultTemplate(ctx context.Context, tx bun.IDB, creatorUuid uuid.UUID, newRecord *entities.ImagingResultTemplateEntity) error {
	newRecord.CreatedAt = time.Now()
	newRecord.CreatedBy = creatorUuid

	return repositories.InsertOne(ctx, tx, newRecord)
}

func UpdateWholeImagingResultTemplateRecord(ctx context.Context, tx bun.IDB, updaterUuid uuid.UUID, template *entities.ImagingResultTemplateEntity) error {
	now := time.Now()
	template.UpdatedAt = &now
	template.UpdatedBy = &updaterUuid

	return repositories.UpdateWholeRecord(ctx, tx, template)
}
