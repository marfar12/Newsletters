package transport

import (
	"database/sql"
	"net/http"

	"newsletter/config"
	"newsletter/db"
	"newsletter/service"
	"newsletter/transport/model"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Port    int
	Mux     *chi.Mux
	Service model.Service
	DB      *sql.DB
}

func Initialize(cfg config.Config) *Handler {
	h := &Handler{
		Port:    cfg.Port,
		Mux:     chi.NewRouter(),
		Service: service.CreateService(),
		DB:      db.Connect(cfg),
	}

	h.Mux.Use(commonMiddleware)

	h.Mux.Route("/newsletter", func(r chi.Router) {
		r.Get("/", h.ListNewsletters)
		r.Post("/", h.CreateNewsletter)

		r.Get("/{id}", h.GetNewsletter)
	})

	return h
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
