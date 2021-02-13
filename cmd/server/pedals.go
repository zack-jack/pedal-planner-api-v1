package main

import (
	"net/http"

	"github.com/zack-jack/pedal-tetris-api-v1/data"
)

func getAllPedals(w http.ResponseWriter, r *http.Request) {
	pedals := data.FindAllPedals()

	jsonResponse(w, r, http.StatusOK, map[string]interface{}{
		"pedals": pedals,
	})
}
