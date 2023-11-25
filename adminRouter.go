package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
)

func createAdminRouter(config *apiConfig) http.Handler{
	adminRouter := chi.NewRouter()

	adminRouter.Get("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/html; charset=utf-8")
		r.Header.Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		html := fmt.Sprintf(`
        <html>
            <body>
                <h1>Welcome, Chirpy Admin</h1>
                <p>Chirpy has been visited %d times!</p>
            </body>
        </html>`, config.fileserverhits)
		w.Write([]byte(html))
	}))
    return adminRouter
}
