package api

import (
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		s.logger.Info(r.Method + " " + r.URL.Path + " started")

		next.ServeHTTP(w, r)

		s.logger.Info(r.Method + " " + r.URL.Path + " completed in " + time.Since(start).String())
	})
}
