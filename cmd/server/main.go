package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/zack-jack/pedal-tetris-api-v1/internal/pedalboards"
	"github.com/zack-jack/pedal-tetris-api-v1/internal/pedals"
)

var (
	// These are use for build time variables
	commit = "1"
	gitTag = "1"

	l = zerolog.New(os.Stderr).With().Timestamp().Logger()
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var cfg struct {
		PedalsDB struct {
			WriterDSN    string `required:"true" envconfig:"PEDALS_WRITER_DSN"`
			MaxOpenConns int    `default:"100" envconfig:"MAX_OPEN_CONNECTIONS"`
		}
		PedalboardsDB struct {
			WriterDSN    string `required:"true" envconfig:"PEDALBOARDS_WRITER_DSN"`
			MaxOpenConns int    `default:"100" envconfig:"MAX_OPEN_CONNECTIONS"`
		}
		Build struct {
			Version string `envconfig:"BUILD_VERSION"`
			Env     string `envconfig:"APP_ENV"`
		}
		Web struct {
			ReadTimeout     time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
			WriteTimeout    time.Duration `default:"30s" envconfig:"WRITE_TIMEOUT"`
			IdleTimeout     time.Duration `default:"120s" envconfig:"IDLE_TIMEOUT"`
			ShutdownTimeout time.Duration `default:"5s" envconfig:"SHUTDOWN_TIMEOUT"`
			Port            string        `default:"5000" envconfig:"PORT"`
		}
	}

	if err := envconfig.Process("", &cfg); err != nil {
		l.Fatal().Err(err).Msg("error parsing config")
	}

	// pedals store
	pedalsStore, err := setupPedalsStore(cfg.PedalsDB.WriterDSN, cfg.PedalsDB.MaxOpenConns)
	if err != nil {
		l.Panic().Err(err).Msg("could not create pedals store")
	}

	// pedalboards store
	pedalboardsStore, err := setupPedalboardsStore(cfg.PedalboardsDB.WriterDSN, cfg.PedalboardsDB.MaxOpenConns)
	if err != nil {
		l.Panic().Err(err).Msg("could not create pedalboards store")
	}

	// create router
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, nil, http.StatusOK, map[string]string{
			"status":     "healthy",
			"version":    gitTag,
			"commitHash": commit,
		})
	}).Methods(http.MethodGet)

	// pedals service
	pedalsSvc, err := pedals.New(pedalsStore)
	if err != nil {
		l.Panic().Err(err).Msg("could not create pedals service")
	}

	// pedalboards service
	pedalboardsSvc, err := pedalboards.New(pedalboardsStore)
	if err != nil {
		l.Panic().Err(err).Msg("could not create pedalboards service")
	}

	// api v1
	v1SubRouter := router.PathPrefix("/v1").Subrouter()
	v1, err := newRouter(v1SubRouter)
	if err != nil {
		l.Panic().Err(err).Msg("could not create router")
	}

	// attach routes
	attachPedalsRoutes(v1, &pedalsHandler{store: pedalsStore, pedalsSvc: pedalsSvc}, l)
	attachPedalboardsRoutes(v1, &pedalboardsHandler{store: pedalboardsStore, pedalboardsSvc: pedalboardsSvc}, l)

	server := &http.Server{
		Addr:         ":" + cfg.Web.Port,
		Handler:      router,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
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
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		if err := server.Shutdown(ctx); err != nil {
			l.Error().Err(err).Msgf("main : Graceful shutdown did not complete in %v", cfg.Web.ShutdownTimeout)

			if err := server.Close(); err != nil {
				l.Fatal().Err(err).Msg("main : Could not stop http server")
			}
		}
	}
}
