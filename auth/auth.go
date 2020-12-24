package auth

import (
	"net/http"
)

func BasicAuthentication(hand http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()

		if !ok || username != "admin" || password != "admin" {
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return
		}

		hand.ServeHTTP(w, req)
	}
}