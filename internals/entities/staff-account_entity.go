package entities

import (
	"time"

	"telerad-core-module/internals/types"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type StaffAccountEntity struct {
	bun.BaseModel `bun:"table:telerad.staff_account"`

	Uuid                  uuid.UUID   `json:"uuid" bun:"uuid,pk,nullzero" db:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt             time.Time   `json:"createdAt" bun:"created_at,notnull,nullzero" db:"created_at" gorm:"column:created_at;not null"`
	CreatedBy             *uuid.UUID  `json:"createdBy" bun:"created_by" db:"created_by" gorm:"column:created_by"`
	UpdatedAt             *time.Time  `json:"updatedAt" bun:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	UpdatedBy             *uuid.UUID  `json:"updatedBy" bun:"updated_by" db:"updated_by" gorm:"column:updated_by"`
	Username              *string     `json:"username" bun:"username" db:"username" gorm:"column:username"`
	PasswordHash          *string     `json:"-" bun:"password_hash" db:"password_hash" gorm:"column:password_hash"`
	Code                  string      `json:"code" bun:"code,notnull" db:"code" gorm:"column:code;not null"`
	FullName              string      `json:"fullName" bun:"full_name,notnull" db:"full_name" gorm:"column:full_name;not null"`
	DateOfBirth           *types.Date `json:"dateOfBirth" bun:"date_of_birth" db:"date_of_birth" gorm:"column:date_of_birth"`
	Gender                string      `json:"gender" bun:"gender,notnull" db:"gender" gorm:"column:gender;not null"`
	CitizenIdentityNumber *string     `json:"citizenIdentityNumber" bun:"citizen_identity_number" db:"citizen_identity_number" gorm:"column:citizen_identity_number"`
	Phone                 *string     `json:"phone" bun:"phone" db:"phone" gorm:"column:phone"`
	Email                 *string     `json:"email" bun:"email" db:"email" gorm:"column:email"`
	FullAddress           *string     `json:"fullAddress" bun:"full_address" db:"full_address" gorm:"column:full_address"`
	IsActive              bool        `json:"isActive" bun:"is_active,notnull" db:"is_active" gorm:"column:is_active;not null"`
	Modalities            []string    `json:"modalities" bun:"modalities,array" db:"modalities" gorm:"column:modalities"`
	Roles                 []string    `json:"roles" bun:"roles,array" db:"roles" gorm:"column:roles"`
	TeleradPartnerUuids   []uuid.UUID `json:"teleradPartnerUuids" bun:"telerad_partner_uuids,array" db:"telerad_partner_uuids" gorm:"column:telerad_partner_uuids"`
}

func (StaffAccountEntity) TableName() string { return "telerad.staff_account" }
