package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Peer Note 👋\n"))
		w.Write([]byte("Server Running Fine\n"))
	})

	return r
}
