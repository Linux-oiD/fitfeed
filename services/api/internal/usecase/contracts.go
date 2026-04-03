package usecase

import (
	"context"
	"fitfeed/api/internal/entity"

	"github.com/google/uuid"
)

type (
	UserManager interface {
		GetProfile(ctx context.Context, username string) (entity.User, error)
		UpdateProfile(ctx context.Context, id uuid.UUID, profile entity.Profile) error
	}
)
