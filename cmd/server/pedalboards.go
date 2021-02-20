package main

import (
	"net/http"

	zlog "github.com/rs/zerolog/log"
	"github.com/zack-jack/pedal-planner-api-v1/internal/pedalboards"
)

type pedalboardsHandler struct {
	store          pedalboards.Store
	pedalboardsSvc *pedalboards.Service
}

type getPedalboardsResponse struct {
	Pedalboards []pedalboards.PedalboardPublic `json:"pedalboards"`
}

func (h *pedalboardsHandler) getAllPedalboards(w http.ResponseWriter, r *http.Request) {
	l := zlog.With().Str("handler", "getAllPedalboards").Logger()

	pedalboards, err := h.pedalboardsSvc.FindAllPedalboards(r.Context())
	if err != nil {
		l.Error().Err(err).Msgf("unable to find pedalboards")
		translateErr(w, r, err)
		return
	}

	jsonResponse(w, r, http.StatusOK, getPedalboardsResponse{
		Pedalboards: pedalboards,
	})
}
