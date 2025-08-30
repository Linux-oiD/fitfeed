package repo

import (
	"context"

	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
)

type (
	UserDB interface {
		Create(context.Context, entity.User) error
		GetByID(context.Context, uuid.UUID) (entity.User, error)
		GetByUsername(context.Context, string) (entity.User, error)
		UpdateUsername(context.Context, uuid.UUID, string) error
		Delete(context.Context, uuid.UUID) error
	}

	ProfileDB interface {
		GetByID(context.Context, uuid.UUID) (entity.Profile, error)
		GetByEmail(context.Context, string) (entity.Profile, error)
		Update(context.Context, uuid.UUID, entity.Profile) error
	}

	OauthDB interface {
		Create(context.Context, entity.OauthProvider) error
		GetByID(context.Context, uuid.UUID) (entity.OauthProvider, error)
		GetByProviderID(context.Context, string) (entity.OauthProvider, error)
		UpdateProviderID(context.Context, uuid.UUID, string) error
		Delete(context.Context, uuid.UUID) error
	}
)
