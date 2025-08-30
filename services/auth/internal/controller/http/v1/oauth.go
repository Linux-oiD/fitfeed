package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (h *V1) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(gothUser)

	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

func (h *V1) getAuthFunction(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println(gothUser)

	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *V1) getLogoutFunction(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	gothic.Logout(w, r)
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}
