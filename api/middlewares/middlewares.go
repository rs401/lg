package middlewares

import "github.com/gorilla/mux"

func SetupMiddleWares(r *mux.Router) {
	r.Use(LogMiddleware)
	r.Use(HeadersMiddleware)
}
