package v1

import (
	"fitfeed/api/internal/config"
	"fitfeed/api/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func NewRouter(r chi.Router, u usecase.UserManager, conf *config.AppConfig) {
	uc := NewUserController(u)
	cc := NewConfigController(conf)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{username}", uc.GetProfile)
		r.Put("/profile", uc.UpdateProfile)
	})

	r.Get("/config", cc.GetConfig)
}
