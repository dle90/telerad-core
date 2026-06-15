package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TeleradPartnerEntity struct {
	bun.BaseModel `bun:"table:telerad.telerad_partner"`

	Uuid            uuid.UUID  `json:"uuid" bun:"uuid,pk,nullzero" db:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt       time.Time  `json:"createdAt" bun:"created_at,notnull,nullzero" db:"created_at" gorm:"column:created_at;not null"`
	CreatedBy       uuid.UUID  `json:"createdBy" bun:"created_by,notnull" db:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt       *time.Time `json:"updatedAt" bun:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	UpdatedBy       *uuid.UUID `json:"updatedBy" bun:"updated_by" db:"updated_by" gorm:"column:updated_by"`
	Code            string     `json:"code" bun:"code,notnull" db:"code" gorm:"column:code;not null"`
	Name            string     `json:"name" bun:"name,notnull" db:"name" gorm:"column:name;not null"`
	IsActive        bool       `json:"isActive" bun:"is_active,notnull" db:"is_active" gorm:"column:is_active;not null"`
	Contact         *string    `json:"contact" bun:"contact" db:"contact" gorm:"column:contact"`
	Callback        bool       `json:"callback" bun:"callback,notnull" db:"callback" gorm:"column:callback;not null"`
	CallbackUrl     *string    `json:"callbackUrl" bun:"callback_url" db:"callback_url" gorm:"column:callback_url"`
	Username        string     `json:"username" bun:"username,notnull" db:"username" gorm:"column:username;not null"`
	PasswordHash    string     `json:"-" bun:"password_hash,notnull" db:"password_hash" gorm:"column:password_hash;not null"`
	PartnerUsername *string    `json:"partnerUsername" bun:"partner_username" db:"partner_username" gorm:"column:partner_username"`
	PartnerPassword *string    `json:"-" bun:"partner_password" db:"partner_password" gorm:"column:partner_password"`
	Modalities      []string   `json:"modalities" bun:"modalities,array" db:"modalities" gorm:"column:modalities"`
}

func (TeleradPartnerEntity) TableName() string { return "telerad.telerad_partner" }
