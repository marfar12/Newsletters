package util

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetIdFromURL(r *http.Request) string {
	id := chi.URLParam(r, "id")

	return id
}
