package http

import (
	v1 "fitfeed/api/internal/controller/http/v1"
	"fitfeed/api/internal/config"
	"fitfeed/api/internal/usecase"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func New(u usecase.UserManager, conf *config.AppConfig) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api root."))
	})

	apiV1 := chi.NewRouter()

	v1.NewRouter(apiV1, u, conf)

	r.Mount("/v1", apiV1)

	return r

}
