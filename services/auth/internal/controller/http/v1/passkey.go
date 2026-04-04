package v1

import (
	"encoding/json"
	"fitfeed/auth/internal/entity"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/markbates/goth/gothic"
)

func (h *V1) beginRegistration(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	user, err := h.u.GetByUsername(r.Context(), username)
	if err != nil {
		user = entity.User{Username: username}
		err = h.u.RegisterUser(r.Context(), &user)
		if err != nil {
			http.Error(w, "failed to register user", http.StatusInternalServerError)
			return
		}
		user, _ = h.u.GetByUsername(r.Context(), username)
	}

	options, sessionData, err := h.pk.BeginRegistration(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := gothic.Store.Get(r, "webauthn-session")
	data, _ := json.Marshal(sessionData)
	session.Values["registration-data"] = string(data)
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

func (h *V1) finishRegistration(w http.ResponseWriter, r *http.Request) {
	slog.Info("check request cookies", "cookies", r.Header.Get("Cookie"))
	username := r.URL.Query().Get("username")
	user, err := h.u.GetByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	session, _ := gothic.Store.Get(r, "webauthn-session")
	sessionDataStr, ok := session.Values["registration-data"].(string)
	if !ok {
		http.Error(w, "session not found", http.StatusBadRequest)
		return
	}

	var sessionData webauthn.SessionData
	json.Unmarshal([]byte(sessionDataStr), &sessionData)

	err = h.pk.FinishRegistration(r.Context(), user, sessionData, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("registration successful"))
}

func (h *V1) beginLogin(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	options, sessionData, err := h.pk.BeginLogin(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := gothic.Store.Get(r, "webauthn-session")
	data, _ := json.Marshal(sessionData)
	session.Values["login-data"] = string(data)
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

func (h *V1) finishLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := gothic.Store.Get(r, "webauthn-session")
	sessionDataStr, ok := session.Values["login-data"].(string)
	if !ok {
		http.Error(w, "session not found", http.StatusBadRequest)
		return
	}

	var sessionData webauthn.SessionData
	json.Unmarshal([]byte(sessionDataStr), &sessionData)

	user, err := h.pk.FinishLogin(r.Context(), sessionData, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := h.j.GenerateToken(user)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Also return user info as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
