package internal

import (
	"log"
	"net/http"
)

func logMidlleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Host: %v, method: %v, path: %v", r.Host, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
