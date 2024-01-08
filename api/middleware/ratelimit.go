package middleware

import (
	"net/http"

	"golang.org/x/sync/semaphore"
)

type RateLimit struct {
	ConnectionsSem *semaphore.Weighted
}

func (rate RateLimit) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok := rate.ConnectionsSem.TryAcquire(1); !ok {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("hit connection limit"))
			return
		}
		defer rate.ConnectionsSem.Release(1)
		next.ServeHTTP(w, r)
	})
}
