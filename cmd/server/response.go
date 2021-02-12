package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, httpStatusCode int, body ...interface{}) {
	if httpStatusCode <= 0 {
		httpStatusCode = http.StatusOK
	}

	if httpStatusCode == http.StatusNoContent {
		// this status must not return a body or a content type
		w.WriteHeader(httpStatusCode)
		return
	}

	if len(body) > 0 {
		// don't write the status code until we set our headers
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(httpStatusCode)

		var data []byte
		data, isAlreadyMarshalled := body[0].([]byte)
		if !isAlreadyMarshalled {
			data, _ = json.Marshal(body[0])
		}
		_, _ = w.Write(data)
		return
	}

	w.WriteHeader(httpStatusCode)
}
