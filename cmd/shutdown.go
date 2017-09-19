package main

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleShutdown ... Captures ctrl-c etc and handles more gracefully.
func HandleShutdown(done chan bool) {
	// Create a channel and listen on it.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to allow capture of Ctrl-C.
	go func() {
		_ = <-sigs
		done <- true
	}()
}
