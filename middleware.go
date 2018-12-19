package main

import (
	"log"
	"net/http"
)

// LogRequestMiddleware literally just log the request URL, and nothing else.
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req = '%s'\n", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
