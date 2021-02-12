package main

import (
	"net/http"
	"os"
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

	// Start the service listening for requests.
	l.Info().Msg("main : API Listening")
	serverErrors <- server.ListenAndServe()
}
