package main

import (
	"crypto/tls"
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
	Router    *mux.Router
	Server    *http.Server
}

// NewServer ... Creates a new server.
func NewServer(port int) (Server, error) {
	if !(port == 80 || port == 443 || port >= 2000) {
		return Server{}, errors.New("the port should be 80, 443, or 2000+")
	}

	return Server{
		Port:      port,
		Hostnames: []string{},
		Router:    mux.NewRouter(),
		Server:    nil,
	}, nil
}

// AddSite ... Adds a new site to the server.
func (s *Server) AddSite(hostname, folder string) {
	s.Hostnames = append(s.Hostnames, hostname)
	sr := s.Router.Host(hostname).Subrouter()
	sr.PathPrefix("/").Handler(s.AddLogging(http.FileServer(http.Dir(folder))))
}

// Serve ... Starts the server going.
func (s *Server) Serve() {
	addr := strconv.Itoa(s.Port)
	fmt.Println()
	fmt.Println("Serving on", addr)
	fmt.Println()

	// Only ONE certificate is used so if you are hosting multiple
	// sites you need a certificate that allows multiple 'common names'.
	// LetsEncrypt certificates support this.
	if addr == "443" {
		// HTTPS

		// Start a standard HTTP server that redirects to HTTPS.
		fmt.Println("Starting HTTP to HTTPS redirect")
		go func() {
			redir := mux.NewRouter()
			redir.HandleFunc("/", s.HTTP2HTTPS)
			plain := &http.Server{
				Handler:      redir,
				Addr:         ":80",
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}
			plain.ListenAndServe()
		}()

		// Start the HTTPS server.
		s.Server = &http.Server{
			Handler:      s.Router,
			Addr:         ":" + addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			TLSConfig:    s.GetTLS(s.Hostnames),
		}
		log.Fatal(s.Server.ListenAndServeTLS("", ""))
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

// GetTLS ... Get TLS support for LetsEncrypt.
func (s *Server) GetTLS(hostnames []string) *tls.Config {
	cache := autocert.DirCache("letsencrypt")
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      cache,
		HostPolicy: autocert.HostWhitelist(hostnames...),
	}
	return &tls.Config{GetCertificate: m.GetCertificate}
}

// HTTP2HTTPS ... Redirects from HTTP to HTTPS.
func (s *Server) HTTP2HTTPS(w http.ResponseWriter, r *http.Request) {
	cs := r.Host
	r.URL.Scheme = "https"
	r.URL.Host = cs
	http.Redirect(w, r, r.URL.String(), http.StatusFound)
}
