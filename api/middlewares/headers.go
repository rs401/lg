// Package middlewares provides middlewares
package middlewares

import (
	"net/http"
)

// HeadersMiddleware sets "Content-Type" header to "application/json"
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
}
