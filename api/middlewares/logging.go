package middlewares

import (
	"log"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			log.Printf(" %s ==> %s %s", r.Method, r.Host, r.URL)
		})
}
