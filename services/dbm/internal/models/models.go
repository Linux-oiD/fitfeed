package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

// User is the model for the user table.
type User struct {
	Base
	Username       string          `gorm:"size:255;index;unique;not null" json:"username"`
	Profile        Profile         `json:"profile"`
	OauthProviders []OauthProvider `json:"oauth_providers"`
}

// Profile is the model for the profile table.
type Profile struct {
	Base
	FirstName string    `gorm:"size:255" json:"first_name"`
	LastName  string    `gorm:"size:255" json:"last_name"`
	AvatarURL string    `gorm:"size:255" json:"avatar_url"`
	Email     string    `gorm:"index;unique" json:"email"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
}

type OauthProvider struct {
	Base
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Provider   string    `gorm:"size:31" json:"provider"`
	ProviderID string    `gorm:"size:255;index" json:"provider_id"`
}
