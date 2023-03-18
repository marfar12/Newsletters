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
	Port              int
	Mux               *chi.Mux
	NewsletterService model.NewsletterService
	EditorService     model.EditorService
	DB                *sql.DB
}

func Initialize(cfg config.Config) *Handler {
	h := &Handler{
		Port:              cfg.Port,
		Mux:               chi.NewRouter(),
		NewsletterService: service.CreateNewsletterService(),
		EditorService:     service.CreateEditorService(),
		DB:                db.Connect(cfg),
	}

	h.Mux.Use(commonMiddleware)

	h.Mux.Route("/newsletter", func(r chi.Router) {
		r.Get("/", h.ListNewsletters)
		r.Post("/", h.CreateNewsletter)

		r.Get("/{id}", h.GetNewsletter)
		r.Patch("/{id}", h.UpdateNewsletter)
		r.Delete("/{id}", h.DeleteNewsletter)
	})

	h.Mux.Route("/auth", func(r chi.Router) {
		r.Post("/signin", h.SignIn)
		r.Post("/signup", h.SignUp)
	})

	return h
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
