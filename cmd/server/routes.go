package main

import (
	"net/http"

	"github.com/rs/zerolog"
)

func attachPedalsRoutes(v1 *router, h *pedalsHandler, lo zerolog.Logger) {
	lo = lo.With().Str("handler", "getAllPedals").Logger()

	v1.HandleFunc("/pedals", h.getAllPedals).Methods(http.MethodGet)
}
