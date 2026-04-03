package usermanager

import (
	"context"
	"fitfeed/api/internal/entity"
	"fitfeed/api/internal/repo"
	"log/slog"

	"github.com/google/uuid"
)

type UserManager struct {
	userDB    repo.UserDB
	profileDB repo.ProfileDB
	logger    *slog.Logger
}

func New(userDB repo.UserDB, profileDB repo.ProfileDB, logger *slog.Logger) *UserManager {
	return &UserManager{userDB: userDB, profileDB: profileDB, logger: logger}
}

func (u *UserManager) GetProfile(ctx context.Context, username string) (entity.User, error) {
	user, err := u.userDB.GetByUsername(ctx, username)
	if err != nil {
		u.logger.Error("failed to get user by username", "error", err, "username", username)
		return entity.User{}, entity.ENOTFOUND
	}

	profile, err := u.profileDB.GetByUserID(ctx, user.ID)
	if err != nil {
		u.logger.Error("failed to get profile by user id", "error", err, "user_id", user.ID)
		return entity.User{}, entity.EINTERNAL
	}

	user.Profile = profile
	return user, nil
}

func (u *UserManager) UpdateProfile(ctx context.Context, id uuid.UUID, profile entity.Profile) error {
	err := u.profileDB.Update(ctx, id, profile)
	if err != nil {
		u.logger.Error("failed to update profile", "error", err, "user_id", id)
		return entity.EINTERNAL
	}
	return nil
}
