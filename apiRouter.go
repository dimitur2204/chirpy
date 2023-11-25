
package main

import (
    "net/http" 
	"github.com/go-chi/chi/v5"
)

func createApiRouter(config *apiConfig) http.Handler{

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	apiRouter.Get("/reset", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		config.fileserverhits = 0
		w.Write([]byte("OK"))
	}))
    apiRouter.Post("/validate_chirp", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.Header.Set("Content-Type", "application/json; charset=utf-8")
        r.Header.Set("Cache-Control", "no-cache")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }))
    return apiRouter
}
