package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
)

// Server ... Handles all sites.
type Server struct {
	Port    int
	Logging bool
	Router  *mux.Router
	Server  *http.Server
}

// NewServer ... Creates a new server.
func NewServer(port int, logging bool) (Server, error) {
	if !(port == 80 || port == 443 || port >= 2000) {
		return Server{}, errors.New("the port should be 80, 443, or 2000+")
	}

	return Server{
		Port:    port,
		Logging: logging,
		Router:  mux.NewRouter(),
		Server:  nil,
	}, nil
}

// AddSite ... Adds a new site to the server.
func (s *Server) AddSite(hostname, folder string) {
	sr := s.Router.Host(hostname).Subrouter()
	sr.PathPrefix("/").Handler(http.FileServer(http.Dir(folder)))
}

// Serve ... Starts the server going.
func (s *Server) Serve() {
	addr := strconv.Itoa(s.Port)
	fmt.Println()
	fmt.Println("Serving on", addr)
	s.Server = &http.Server{
		Handler:      s.Router,
		Addr:         ":" + addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.Server.ListenAndServe())
}
