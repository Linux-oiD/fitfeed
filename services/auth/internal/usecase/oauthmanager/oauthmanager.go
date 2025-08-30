package oauthmanager

import (
	"context"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"

	"github.com/google/uuid"
)

type UseCase struct {
	db repo.OauthDB
}

// Add new Oauth provider to the User
func (u *UseCase) AddProvider(ctx context.Context, provider entity.OauthProvider) error {
	panic("not implemented")
}

// Update OauthProvider object. Returns EUNAUTHORIZED if current user is not
// the owner of provider that is being updated. Returns ENOTFOUND if provider does not exist.
func (u *UseCase) UpdateProviderID(ctx context.Context, id uuid.UUID, providerID string) {
	panic("not implemented") // TODO: Implement
}

// Delete provider object
func (u *UseCase) DeleteProvider(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func New(db repo.OauthDB) *UseCase {
	return &UseCase{db: db}
}
