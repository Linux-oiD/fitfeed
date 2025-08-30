package usermanager

import (
	"context"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"

	"github.com/google/uuid"
)

type UserManager struct {
	db repo.UserDB
}

func New(db repo.UserDB) *UserManager {

	return &UserManager{db: db}
}

// Check username availability. If username is not available -
// returns ENOTAVAILABLE
func (u *UserManager) CheckUsername(ctx context.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Create new user.
func (u *UserManager) RegisterUser(ctx context.Context, user entity.User) error {
	panic("not implemented") // TODO: Implement
}

// Updates a user object. Returns EUNAUTHORIZED if current user is not
// the user that is being updated. Returns ENOTFOUND if user does not exist.
func (u *UserManager) UpdateUsername(ctx context.Context, id uuid.UUID, username string) (entity.User, error) {
	panic("not implemented") // TODO: Implement
}

// Permanently deletes a user and all owned dials. Returns EUNAUTHORIZED
// if current user is not the user being deleted. Returns ENOTFOUND if
// user does not exist.
func (u *UserManager) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
