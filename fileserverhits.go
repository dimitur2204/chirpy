package main

import "net/http"

func (cfg *apiConfig) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Cache-Control", "no-cache")
		cfg.fileserverhits++
		next.ServeHTTP(w, r)
	})
}
