package oauthmanager

import (
	"context"
	"errors"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UseCase struct {
	db     repo.OauthDB
	logger *slog.Logger
}

func New(db repo.OauthDB, logger *slog.Logger) *UseCase {
	return &UseCase{db: db, logger: logger}
}

// Add new Oauth provider to the User
func (u *UseCase) AddProvider(ctx context.Context, provider entity.OauthProvider) error {
	err := u.db.Create(ctx, provider)
	if err != nil {
		u.logger.Error("failed to create oauth provider", "error", err, "provider", provider.Provider, "user_id", provider.UserID)
		return entity.EINTERNAL
	}
	return nil
}

func (u *UseCase) GetByProviderID(ctx context.Context, providerID string) (entity.OauthProvider, error) {
	provider, err := u.db.GetByProviderID(ctx, providerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.OauthProvider{}, entity.ENOTFOUND
		}
		u.logger.Error("failed to get provider by provider id", "error", err, "provider_id", providerID)
		return entity.OauthProvider{}, entity.EINTERNAL
	}
	return provider, nil
}

// Update OauthProvider object. Returns EUNAUTHORIZED if current user is not
// the owner of provider that is being updated. Returns ENOTFOUND if provider does not exist.
func (u *UseCase) UpdateProviderID(ctx context.Context, id uuid.UUID, providerID string) {
	err := u.db.UpdateProviderID(ctx, id, providerID)
	if err != nil {
		u.logger.Error("failed to update provider id", "error", err, "id", id, "provider_id", providerID)
	}
}

// Delete provider object
func (u *UseCase) DeleteProvider(ctx context.Context, id uuid.UUID) error {
	err := u.db.Delete(ctx, id)
	if err != nil {
		u.logger.Error("failed to delete provider", "error", err, "id", id)
		return entity.EINTERNAL
	}
	return nil
}
