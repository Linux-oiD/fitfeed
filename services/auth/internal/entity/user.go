package entity

import (
	"context"

	"github.com/google/uuid"
)

// User is the model for the user table.
type User struct {
	Base
	Username       string          `gorm:"size:255;uniqueIndex;not null" json:"username"`
	Profile        Profile         `json:"profile"`
	OauthProviders []OauthProvider `json:"oauth_providers"`
}

type UserService interface {

	// Retrieves a user by ID along with their associated auth objects.
	// Returns ENOTFOUND if user does not exist.
	FindUserByID(ctx context.Context, id uuid.UUID) (*User, error)

	// Retrieves a list of users by filter. Also returns total count of matching
	// users which may differ from returned results if filter.Limit is specified.
	FindUsers(ctx context.Context, filter UserFilter) ([]*User, int, error)
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	// Filtering fields.
	ID       *uuid.UUID `json:"id"`
	Username *string    `json:"username"`

	// Restrict to subset of results.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
