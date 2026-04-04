package usermanager

import (
	"context"
	"errors"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserManager struct {
	db        repo.UserDB
	profileDB repo.ProfileDB
	logger    *slog.Logger
}

func New(db repo.UserDB, profileDB repo.ProfileDB, logger *slog.Logger) *UserManager {
	return &UserManager{db: db, profileDB: profileDB, logger: logger}
}

func (u *UserManager) CheckUsername(ctx context.Context, username string) error {
	_, err := u.db.GetByUsername(ctx, username)
	if err == nil {
		return entity.ENOTAVAILABLE
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	u.logger.Error("failed to check username", "error", err, "username", username)
	return entity.EINTERNAL
}

func (u *UserManager) RegisterUser(ctx context.Context, user *entity.User) error {
	err := u.db.Create(ctx, user)
	if err != nil {
		u.logger.Error("failed to register user", "error", err, "username", user.Username)
		return entity.EINTERNAL
	}

	// Also create profile
	user.Profile.UserID = user.ID
	err = u.profileDB.Create(ctx, user.Profile)
	if err != nil {
		u.logger.Error("failed to create profile for user", "error", err, "username", user.Username)
		// NOTE: In a real system we should use a transaction here.
		return entity.EINTERNAL
	}

	return nil
}

func (u *UserManager) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	user, err := u.db.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, entity.ENOTFOUND
		}
		u.logger.Error("failed to get user by username", "error", err, "username", username)
		return entity.User{}, entity.EINTERNAL
	}
	return user, nil
}

func (u *UserManager) UpdateUsername(ctx context.Context, id uuid.UUID, username string) (entity.User, error) {
	err := u.db.UpdateUsername(ctx, id, username)
	if err != nil {
		u.logger.Error("failed to update username", "error", err, "id", id, "username", username)
		return entity.User{}, entity.EINTERNAL
	}
	return u.db.GetByID(ctx, id)
}

func (u *UserManager) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := u.db.Delete(ctx, id)
	if err != nil {
		u.logger.Error("failed to delete user", "error", err, "id", id)
		return entity.EINTERNAL
	}
	return nil
}
