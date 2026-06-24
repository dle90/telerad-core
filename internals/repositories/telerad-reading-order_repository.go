package repositories

import (
	"context"
	"time"

	"telerad-core-module/internals/entities"
	databaseQueryModels "telerad-core-module/internals/models/database-query_models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
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

// FindPaginatedReadingOrders trả danh sách ca đọc (kèm tên đối tác + tên bác sĩ
// đọc) đã lọc/scope, sắp xếp ngày chụp mới nhất trước.
func FindPaginatedReadingOrders(
	ctx context.Context,
	tx bun.IDB,
	page, pageSize int,
	filter ReadingOrderListFilter,
) ([]databaseQueryModels.ReadingOrderListRow, int, error) {
	var rows []databaseQueryModels.ReadingOrderListRow

	query := tx.NewSelect().
		Model(&rows).
		ModelTableExpr("telerad.telerad_reading_order AS ro").
		ColumnExpr("ro.uuid").
		ColumnExpr("ro.telerad_partner_uuid").
		ColumnExpr("ro.order_code").
		ColumnExpr("ro.order_item_code").
		ColumnExpr("ro.study_instance_uid").
		ColumnExpr("ro.patient_code").
		ColumnExpr("ro.full_name").
		ColumnExpr("ro.gender").
		ColumnExpr("ro.years_old").
		ColumnExpr("ro.phone").
		ColumnExpr("ro.service_name").
		ColumnExpr("ro.modality").
		ColumnExpr("ro.modality_name").
		ColumnExpr("ro.perform_ended_at").
		ColumnExpr("ro.read_completed_at").
		ColumnExpr("ro.assigned_to").
		ColumnExpr("ro.status").
		ColumnExpr("ro.result_returned").
		ColumnExpr("tp.name AS partner_name").
		ColumnExpr("sa.full_name AS assigned_to_name").
		Join("JOIN telerad.telerad_partner AS tp ON tp.uuid = ro.telerad_partner_uuid").
		Join("LEFT JOIN telerad.staff_account AS sa ON sa.uuid = ro.assigned_to")

	if !filter.IsAdmin { // Nếu không phải admin
		if len(filter.PartnerUuids) == 0 || len(filter.Modalities) == 0 {
			return rows, 0, nil
		}

		query = query.Where("ro.telerad_partner_uuid IN (?)", bun.List(filter.PartnerUuids)).
			Where("ro.modality IN (?)", bun.List(filter.Modalities))
	} else {
		if len(filter.PartnerUuids) > 0 {
			query = query.Where("ro.telerad_partner_uuid IN (?)", bun.List(filter.PartnerUuids))
		}
		if len(filter.Modalities) > 0 {
			query = query.Where("ro.modality IN (?)", bun.List(filter.Modalities))
		}
	}

	// Lọc theo ngày chụp (perform_ended_at).
	if filter.PerformEndedFrom != nil {
		query = query.Where("ro.perform_ended_at >= ?", *filter.PerformEndedFrom)
	}
	if filter.PerformEndedTo != nil {
		query = query.Where("ro.perform_ended_at <= ?", *filter.PerformEndedTo)
	}

	// Lọc text.
	if filter.PatientName != "" {
		query = query.Where("ro.full_name ILIKE ?", "%"+filter.PatientName+"%")
	}
	if filter.PatientCode != "" {
		query = query.Where("ro.patient_code ILIKE ?", "%"+filter.PatientCode+"%")
	}
	if filter.Phone != "" {
		query = query.Where("ro.phone = ?", filter.Phone)
	}

	// Lọc theo tình trạng ca + bác sĩ đang nhận + trạng thái trả kết quả.
	if filter.Status != "" {
		query = query.Where("ro.status = ?", filter.Status)
	}
	if filter.AssignedTo != nil {
		query = query.Where("ro.assigned_to = ?", *filter.AssignedTo)
	}
	if filter.ResultReturned != nil {
		query = query.Where("ro.result_returned = ?", *filter.ResultReturned)
	}

	query = query.OrderExpr("ro.perform_ended_at DESC, ro.uuid ASC")

	totalCount, err := findPaginated(ctx, query, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return rows, totalCount, nil
}

// FindOneReadingOrderByPartnerAndOrderItemId dùng để chống trùng: 1 đối tác không
// được đẩy lặp cùng một order_item_id.
func FindOneReadingOrderByPartnerAndOrderItemId(ctx context.Context, tx bun.IDB, teleradPartnerUuid uuid.UUID, orderItemId string) (*entities.TeleradReadingOrderEntity, error) {
	var record entities.TeleradReadingOrderEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.telerad_partner_uuid = ?", teleradPartnerUuid).
		Where("?TableAlias.order_item_id = ?", orderItemId)

	return findOne(ctx, query, &record)
}

// FindOneReadingOrderByAssigneeAndStatus tìm 1 ca đọc đang gán cho 1 bác sĩ ở 1
// trạng thái (vd kiểm tra user có ca nào đang READING không). Trả nil nếu không có.
func FindOneReadingOrderByAssigneeAndStatus(ctx context.Context, tx bun.IDB, assigneeUuid uuid.UUID, status string) (*entities.TeleradReadingOrderEntity, error) {
	var record entities.TeleradReadingOrderEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.assigned_to = ?", assigneeUuid).
		Where("?TableAlias.status = ?", status).
		Limit(1)

	return findOne(ctx, query, &record)
}
