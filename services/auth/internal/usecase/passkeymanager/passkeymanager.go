package passkeymanager

import (
	"context"
	"fitfeed/auth/internal/entity"
	"fitfeed/auth/internal/repo"
	"log/slog"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type UseCase struct {
	w      *webauthn.WebAuthn
	db     repo.PasskeyDB
	u      repo.UserDB
	logger *slog.Logger
}

func New(w *webauthn.WebAuthn, db repo.PasskeyDB, u repo.UserDB, logger *slog.Logger) *UseCase {
	return &UseCase{w: w, db: db, u: u, logger: logger}
}

func (u *UseCase) BeginRegistration(ctx context.Context, user entity.User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	creation, sessionData, err := u.w.BeginRegistration(user)
	if err != nil {
		u.logger.Error("failed to begin passkey registration", "error", err, "username", user.Username)
		return nil, nil, entity.EINTERNAL
	}
	return creation, sessionData, nil
}

func (u *UseCase) FinishRegistration(ctx context.Context, user entity.User, session webauthn.SessionData, response *http.Request) error {
	credential, err := u.w.FinishRegistration(user, session, response)
	if err != nil {
		u.logger.Error("failed to finish passkey registration", "error", err, "username", user.Username)
		return entity.EINVALID
	}

	pk := entity.Passkey{
		UserID:          user.ID,
		CredentialID:    credential.ID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
		AAGUID:          credential.Authenticator.AAGUID,
		SignCount:       credential.Authenticator.SignCount,
	}

	err = u.db.Create(ctx, pk)
	if err != nil {
		u.logger.Error("failed to save passkey", "error", err, "username", user.Username)
		return entity.EINTERNAL
	}

	return nil
}

func (u *UseCase) BeginLogin(ctx context.Context, username string) (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	user, err := u.u.GetByUsername(ctx, username)
	if err != nil {
		return nil, nil, entity.ENOTFOUND
	}

	assertion, sessionData, err := u.w.BeginLogin(user)
	if err != nil {
		u.logger.Error("failed to begin passkey login", "error", err, "username", username)
		return nil, nil, entity.EINTERNAL
	}
	return assertion, sessionData, nil
}

func (u *UseCase) FinishLogin(ctx context.Context, session webauthn.SessionData, response *http.Request) (entity.User, error) {
	// WebAuthn FinishLogin needs the user, but we don't know who it is yet 
	// until we parse the response or we use the username from sessionData.
	// Actually, sessionData has UserID.
	
	username := string(session.UserID)
	user, err := u.u.GetByUsername(ctx, username)
	if err != nil {
		return entity.User{}, entity.ENOTFOUND
	}

	credential, err := u.w.FinishLogin(user, session, response)
	if err != nil {
		u.logger.Error("failed to finish passkey login", "error", err, "username", username)
		return entity.User{}, entity.EINVALID
	}

	// Update sign count
	err = u.db.UpdateSignCount(ctx, credential.ID, credential.Authenticator.SignCount)
	if err != nil {
		u.logger.Error("failed to update passkey sign count", "error", err, "username", username)
	}

	return user, nil
}
