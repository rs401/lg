package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs401/lg/api/tokenutils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			at, err := tokenutils.VerifyToken(r)
			if err != nil {
				log.Println("Bad token")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			// Check still valid
			if !tokenutils.StillValid(at) {
				// Token expired, check the refresh token
				rt, err := tokenutils.VerifyRefreshToken(r)
				if err != nil {
					log.Println("Bad refresh token")
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
					return
				}
				// Check refresh token still valid
				if !tokenutils.StillValid(rt) {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(map[string]string{"error": "token expired"})
					return
				}
				// Refresh still valid, get them a new token
				newToken, err := tokenutils.RefreshAccessToken(rt)
				if err != nil {
					log.Println("Really Bad refresh token")
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
					return
				}
				// Set new token
				w.Header().Set("Authorization", newToken)
			}

			next.ServeHTTP(w, r)
		})
}
