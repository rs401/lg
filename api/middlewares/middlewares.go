package middlewares

// Package middlewares provides middlewares
import "github.com/gorilla/mux"

// SetupMiddleWares takes a *mux.Router and sets it to use the middlewares
func SetupMiddleWares(r *mux.Router) {
	r.Use(LogMiddleware)
	r.Use(HeadersMiddleware)
}
