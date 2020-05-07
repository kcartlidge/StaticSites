package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

// Server ... Handles all sites.
type Server struct {
	Port      int
	Hostnames []string
	Logging   bool
	Router    *mux.Router
	Server    *http.Server
}

// NewServer ... Creates a new server.
func NewServer(port int, logging bool) (Server, error) {
	if !(port == 80 || port == 443 || port >= 2000) {
		return Server{}, errors.New("the port should be 80, 443, or 2000+")
	}

	return Server{
		Port:      port,
		Hostnames: []string{},
		Logging:   logging,
		Router:    mux.NewRouter(),
		Server:    nil,
	}, nil
}

// AddSite ... Adds a new site to the server.
func (s *Server) AddSite(hostname, folder string) {
	s.Hostnames = append(s.Hostnames, hostname)
	sr := s.Router.Host(hostname).Subrouter()
	if s.Logging {
		sr.PathPrefix("/").Handler(s.AddLogging(http.FileServer(http.Dir(folder))))
	} else {
		sr.PathPrefix("/").Handler(http.FileServer(http.Dir(folder)))
	}
}

// Serve ... Starts the server going.
func (s *Server) Serve() {
	addr := strconv.Itoa(s.Port)
	fmt.Println()
	fmt.Println("Serving on", addr)
	fmt.Println()

	// Only ONE certificate is used so if you are hosting multiple
	// sites you need a certificate that allows multiple common names.
	// LetsEncrypt certificates support this.
	//
	// https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20200429183012-4b2356b1ed79/acme/autocert?tab=doc#NewListener
	// Certificates are cached in a "golang-autocert" directory under an operating
	// system-specific cache or temp directory. This may not be suitable for servers
	// spanning multiple machines.
	if addr == "443" {
		// HTTPS
		log.Fatal(http.Serve(autocert.NewListener(s.Hostnames...), s.Router))
	} else {
		// HTTP
		s.Server = &http.Server{
			Handler:      s.Router,
			Addr:         ":" + addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(s.Server.ListenAndServe())
	}
}
