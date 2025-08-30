package profilemanager

import (
	"context"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"

	"github.com/google/uuid"
)

type UseCase struct {
	db repo.ProfileDB
}

// Update Profile object. Returns EUNAUTHORIZED if current user is not
// the owner of Profile that is being updated. Returns ENOTFOUND if profile does not exist.
func (u *UseCase) UpdateProfile(ctx context.Context, profileUpd entity.ProfileUpdate) (entity.Profile, error) {
	panic("not implemented") // TODO: Implement
}

// Return UserID if email exist overvise return nil and ENOTFOUND if email does not exist.
func (u *UseCase) CheckEmail(ctx context.Context, email string) (uuid.UUID, error) {
	panic("not implemented") // TODO: Implement
}

func New(db repo.ProfileDB) *UseCase {

	return &UseCase{db: db}
}
