package middleware

import (
	"log"
	"net/http"
)

func TokenParse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			log.Printf("no token\n")
		}

		next.ServeHTTP(w, r)
	})
}
