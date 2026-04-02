package v1

import (
	"encoding/json"
	"fitfeed/api/internal/entity"
	"fitfeed/api/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	u usecase.UserManager
}

func NewUserController(u usecase.UserManager) *UserController {
	return &UserController{u: u}
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := c.u.GetProfile(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user ID from JWT context (needs middleware)
	// For now assume we get it from request or similar (unsafe)
	var profile entity.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// err := c.u.UpdateProfile(r.Context(), userID, profile)
	w.WriteHeader(http.StatusNotImplemented)
}
