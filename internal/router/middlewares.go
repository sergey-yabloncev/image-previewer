package router

import (
	"log"
	"net/http"
)

// LoggerMiddleware check extension in request.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
