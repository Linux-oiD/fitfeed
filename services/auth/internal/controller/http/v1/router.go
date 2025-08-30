package v1

import (
	"fitfeed/auth/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func NewUserRoutes(r *chi.Mux, u usecase.UserManager, o usecase.OauthManager, p usecase.ProfileManager) {

	h := &V1{u: u, o: o, p: p}

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.listUsers) // GET /users
		r.Get("/{id}", h.getUserByID)
	})
}

func NewOauthRoutes(r *chi.Mux, u usecase.UserManager, o usecase.OauthManager, p usecase.ProfileManager) {

	h := &V1{u: u, o: o, p: p}

	r.Route("/oauth", func(r chi.Router) {
		r.Get("/{provider}/callback", h.getAuthCallbackFunction)
		r.Get("/{provider}/auth", h.getAuthFunction)
		r.Get("/{provider}/logout", h.getLogoutFunction)

	})
}
