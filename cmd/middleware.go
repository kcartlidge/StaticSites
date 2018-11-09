package main

import (
	"log"
	"net/http"
)

// AddLogging ... Adds logging middleware (to the command console).
func (s *Server) AddLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
