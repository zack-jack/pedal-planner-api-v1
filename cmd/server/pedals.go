package main

import (
	"net/http"

	zlog "github.com/rs/zerolog/log"
	"github.com/zack-jack/pedal-tetris-api-v1/internal/pedals"
)

type pedalsHandler struct {
	store     pedals.Store
	pedalsSvc *pedals.Service
}

type getPedalsResponse struct {
	Pedals []pedals.PedalPublic `json:"pedals"`
}

func (h *pedalsHandler) getAllPedals(w http.ResponseWriter, r *http.Request) {
	l := zlog.With().Str("handler", "getAllPedals").Logger()

	pedals, err := h.pedalsSvc.FindAllPedals(r.Context())
	if err != nil {
		l.Error().Err(err).Msgf("unable to find pedals")
		translateErr(w, r, err)
		return
	}

	jsonResponse(w, r, http.StatusOK, getPedalsResponse{
		Pedals: pedals,
	})
}
