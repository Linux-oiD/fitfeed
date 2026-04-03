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
	claims, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var profile entity.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.u.UpdateProfile(r.Context(), claims.ID, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile updated"})
}
