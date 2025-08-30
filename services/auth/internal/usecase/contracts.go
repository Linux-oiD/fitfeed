package usecase

import (
	"context"

	"fitfeed/auth/internal/entity"

	"github.com/google/uuid"
)

type (
	UserManager interface {
		// Check username availability. If username is not available -
		// returns ENOTAVAILABLE
		CheckUsername(ctx context.Context, username string) error

		//Create new user.
		RegisterUser(ctx context.Context, user entity.User) error

		// Updates a user object. Returns EUNAUTHORIZED if current user is not
		// the user that is being updated. Returns ENOTFOUND if user does not exist.
		UpdateUsername(ctx context.Context, id uuid.UUID, username string) (entity.User, error)

		// Permanently deletes a user and all owned dials. Returns EUNAUTHORIZED
		// if current user is not the user being deleted. Returns ENOTFOUND if
		// user does not exist.
		DeleteUser(ctx context.Context, id uuid.UUID) error
	}

	OauthManager interface {
		//Add new Oauth provider to the User
		AddProvider(ctx context.Context, provider entity.OauthProvider) error

		//Update OauthProvider object. Returns EUNAUTHORIZED if current user is not
		// the owner of provider that is being updated. Returns ENOTFOUND if provider does not exist.
		UpdateProviderID(ctx context.Context, id uuid.UUID, providerID string)

		//Delete provider object
		DeleteProvider(ctx context.Context, id uuid.UUID) error
	}

	ProfileManager interface {
		//Update Profile object. Returns EUNAUTHORIZED if current user is not
		// the owner of Profile that is being updated. Returns ENOTFOUND if profile does not exist.
		UpdateProfile(ctx context.Context, profileUpd entity.ProfileUpdate) (entity.Profile, error)

		// Return UserID if email exist overvise return nil and ENOTFOUND if email does not exist.
		CheckEmail(ctx context.Context, email string) (uuid.UUID, error)
	}
)
