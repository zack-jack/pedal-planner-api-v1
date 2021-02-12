package main

import (
	"net/http"

	"github.com/rs/zerolog"
)

func attachPedalsRoutes(v1 *router, lo zerolog.Logger) {
	lo = lo.With().Str("handler", "pedals").Logger()

	v1.HandleFunc("/pedals", getPedals).Methods(http.MethodGet)
}
