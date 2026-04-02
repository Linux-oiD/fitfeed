package profilemanager

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
	db     repo.ProfileDB
	logger *slog.Logger
}

func New(db repo.ProfileDB, logger *slog.Logger) *UseCase {
	return &UseCase{db: db, logger: logger}
}

// Update Profile object. Returns EUNAUTHORIZED if current user is not
// the owner of Profile that is being updated. Returns ENOTFOUND if profile does not exist.
func (u *UseCase) UpdateProfile(ctx context.Context, profileUpd entity.ProfileUpdate) (entity.Profile, error) {
	// Need to know WHO is updating to check EUNAUTHORIZED. 
	// For now just implementing basic update.
	panic("not fully implemented - needs auth context")
}

// Return UserID if email exist overvise return nil and ENOTFOUND if email does not exist.
func (u *UseCase) CheckEmail(ctx context.Context, email string) (uuid.UUID, error) {
	profile, err := u.db.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, entity.ENOTFOUND
		}
		u.logger.Error("failed to check email", "error", err, "email", email)
		return uuid.Nil, entity.EINTERNAL
	}
	return profile.UserID, nil
}
