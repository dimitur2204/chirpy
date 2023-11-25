package main

import (
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
	corsMux := middlewareCors(r)
	r.Get("/app/*", config.middlewareMetrics(http.StripPrefix("/app", http.FileServer(http.Dir(".")))).(http.HandlerFunc))
	r.Get("/app", config.middlewareMetrics(http.StripPrefix("/app", http.FileServer(http.Dir(".")))).(http.HandlerFunc))
	r.Mount("/api", createApiRouter(&config))
	r.Mount("/admin", createAdminRouter(&config))
	server := http.Server{
		Handler: corsMux,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
