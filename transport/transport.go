package transport

import (
	"database/sql"
	"net/http"

	"newsletter/config"
	"newsletter/db"
	"newsletter/service"
	"newsletter/transport/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	Port              int
	Mux               *chi.Mux
	NewsletterService model.NewsletterService
	DB                *sql.DB
}

func Initialize(cfg config.Config) *Handler {
	h := &Handler{
		Port:              cfg.Port,
		Mux:               chi.NewRouter(),
		NewsletterService: service.CreateNewsletterService(),
		DB:                db.Connect(),
	}

	h.Mux.Use(setContentType)

	h.Mux.Route("/newsletter", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(model.TokenAuth))
			r.Use(jwtauth.Authenticator)

			r.Post("/", h.CreateNewsletter)
			r.Patch("/{id}", h.UpdateNewsletter)
			r.Delete("/{id}", h.DeleteNewsletter)

		})

		r.Get("/", h.ListNewsletters)

		r.Get("/{id}", h.GetNewsletter)

	})

	h.Mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(model.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/publish", h.Publish)
	})

	h.Mux.Route("/auth", func(r chi.Router) {
		r.Post("/signin", h.SignIn)
		r.Post("/signup", h.SignUp)
	})

	h.Mux.Post("/subscribe", h.Subscribe)
	h.Mux.Get("/unsubscribe/{id}", h.Unsubscribe)

	return h
}

func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
