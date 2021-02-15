package main

import (
	"net/http"

	"github.com/rs/zerolog"
)

func attachPedalsRoutes(v1 *router, h *pedalsHandler, lo zerolog.Logger) {
	lo = lo.With().Str("handler", "getAllPedals").Logger()

	v1.HandleFunc("/pedals", h.getAllPedals).Methods(http.MethodGet)
}

func attachPedalboardsRoutes(v1 *router, h *pedalboardsHandler, lo zerolog.Logger) {
	lo = lo.With().Str("handler", "getAllPedalboards").Logger()

	v1.HandleFunc("/pedalboards", h.getAllPedalboards).Methods(http.MethodGet)
}
