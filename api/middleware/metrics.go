package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
}

func WriteMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		before := time.Now()
		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		diff := time.Since(before)
		tDiff := float64(diff / time.Millisecond)

		log.Printf("*** mectrics %s %d %v ***\n", r.Method, rw.status, tDiff)
	})
}
