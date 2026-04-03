package entity

import "github.com/google/uuid"

type OauthProvider struct {
	Base
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Provider   string    `gorm:"size:31" json:"provider"`
	ProviderID string    `gorm:"size:255;index" json:"provider_id"`
}
