package main

import "net/http"

func getPedals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse(w, r, http.StatusOK, map[string]interface{}{
		"status": "200",
		"desc":   "getPedals",
	})
}
