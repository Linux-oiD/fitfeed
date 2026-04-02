package v1

import (
	"context"
	"errors"
	"fitfeed/auth/internal/entity"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (h *V1) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	var user entity.User

	// 1. Check if this OAuth provider is already linked
	oauthProvider, err := h.o.GetByProviderID(ctx, gothUser.UserID)
	if err == nil {
		// Found user by OAuth provider
		user, err = h.u.GetByUsername(ctx, gothUser.NickName) // Assuming NickName is username
		if err != nil && !errors.Is(err, entity.ENOTFOUND) {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		// If user not found by username but oauth exists, we might have a data inconsistency 
		// or we should use UserID from oauthProvider
		// For now, let's assume we find them.
	} else if errors.Is(err, entity.ENOTFOUND) {
		// 2. Provider not found. Check if user with this username exists
		user, err = h.u.GetByUsername(ctx, gothUser.NickName)
		if err == nil {
			// User exists, bind new provider
			err = h.o.AddProvider(ctx, entity.OauthProvider{
				UserID:     user.ID,
				Provider:   provider,
				ProviderID: gothUser.UserID,
			})
			if err != nil {
				http.Error(w, "failed to bind provider", http.StatusInternalServerError)
				return
			}
		} else if errors.Is(err, entity.ENOTFOUND) {
			// 3. New user - Register
			user = entity.User{
				Username: gothUser.NickName,
				Profile: entity.Profile{
					FirstName: gothUser.FirstName,
					LastName:  gothUser.LastName,
					AvatarURL: gothUser.AvatarURL,
					Email:     gothUser.Email,
				},
				OauthProviders: []entity.OauthProvider{
					{
						Provider:   provider,
						ProviderID: gothUser.UserID,
					},
				},
			}
			err = h.u.RegisterUser(ctx, user)
			if err != nil {
				http.Error(w, "failed to register user", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// 4. Generate JWT
	token, err := h.j.GenerateToken(user)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	// 5. Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in prod
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

func (h *V1) getAuthFunction(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println(gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *V1) getLogoutFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.Logout(w, r)

	// Clear JWT cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}
