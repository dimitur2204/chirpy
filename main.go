package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config := apiConfig{
		fileserverhits: 0,
	}
	r := chi.NewRouter()
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	r.Get("/app/*", config.middlewareMetrics(http.StripPrefix("/app", http.FileServer(http.Dir(".")))).(http.HandlerFunc))
	r.Get("/app", config.middlewareMetrics(http.StripPrefix("/app", http.FileServer(http.Dir(".")))).(http.HandlerFunc))
	r.Get("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	r.Get("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits: " + fmt.Sprint(config.fileserverhits)))
	}))
	r.Get("/reset", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		config.fileserverhits = 0
		w.Write([]byte("OK"))
	}))
	server := http.Server{
		Handler: corsMux,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
