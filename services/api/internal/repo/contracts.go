package repo

import (
	"context"
	"fitfeed/api/internal/entity"

	"github.com/google/uuid"
)

type (
	UserDB interface {
		GetByID(context.Context, uuid.UUID) (entity.User, error)
		GetByUsername(context.Context, string) (entity.User, error)
	}

	ProfileDB interface {
		GetByUserID(context.Context, uuid.UUID) (entity.Profile, error)
		Update(context.Context, uuid.UUID, entity.Profile) error
	}
)
