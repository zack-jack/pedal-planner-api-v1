package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var (
	// These are use for build time variables
	commit = "1"
	gitTag = "1"

	l = zerolog.New(os.Stderr).With().Timestamp().Logger()
)

func main() {
	// create router
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, nil, http.StatusOK, map[string]string{
			"status":     "healthy",
			"version":    gitTag,
			"commitHash": commit,
		})
	}).Methods(http.MethodGet)

	// api v1
	v1SubRouter := router.PathPrefix("/v1").Subrouter()
	v1, err := newRouter(v1SubRouter)
	if err != nil {
		l.Panic().Err(err).Msg("could not create router")
	}

	// attach routes
	attachPedalsRoutes(v1, l)

	server := &http.Server{
		Addr:         ":" + "5000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// =========================================================================
	// Start

	// Start the service listening for requests.
	go func() {
		l.Info().Msg("main : API Listening")
		serverErrors <- server.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Stop API Service

	select {
	case err := <-serverErrors:
		l.Fatal().Err(err).Msg("error starting server")

	case <-osSignals:
		l.Info().Msg("main : Start shutdown...")

		// Create context for Shutdown call
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Asking listener to shutdown and load shed.
		if err := server.Shutdown(ctx); err != nil {
			l.Error().Err(err).Msgf("main : Graceful shutdown did not complete in %v", 5*time.Second)

			if err := server.Close(); err != nil {
				l.Fatal().Err(err).Msg("main : Could not stop http server")
			}
		}
	}
}
