package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sahil/peernote/internal/handlers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(api chi.Router) {
		api.Post("/files", handlers.UploadFileMetadata)
		api.Get("/peers", handlers.GetPeersByFileName)
		api.Post("/register", handlers.RegisterPeer)
		api.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Central server is healthy."))
		})
	})

	return r
}
