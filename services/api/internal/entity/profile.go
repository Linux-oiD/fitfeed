package entity

import "github.com/google/uuid"

// Profile is the model for the profile table.
type Profile struct {
	Base
	FirstName string    `gorm:"size:255" json:"first_name"`
	LastName  string    `gorm:"size:255" json:"last_name"`
	AvatarURL string    `gorm:"size:255" json:"avatar_url"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
}

type ProfileUpdate struct {
	FirsrName string
	LastName  string
	AvatarURL string
	Email     string
}
