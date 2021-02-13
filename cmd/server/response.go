package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	api "github.com/zack-jack/pedal-tetris-api-v1"
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

func translateErr(w http.ResponseWriter, r *http.Request, err error, invalidParam ...string) {
	jsonMsgResponse := func(wr http.ResponseWriter, re *http.Request, code, msg string, httpStatusCode int) {
		jsonResponse(wr, re, httpStatusCode, map[string]string{
			"code":    code,
			"message": msg,
		})
	}

	switch errors.Cause(err) {
	case api.ErrCouldNotReadBody:
		jsonMsgResponse(w, r, "0000", "could not read body", http.StatusBadRequest)
	case api.ErrInvalidParam:
		// valid lengths for invalidParam are essentially 0, 1, or 2
		switch len(invalidParam) {
		case 0:
			jsonMsgResponse(w, r, "0003", "invalid parameter", http.StatusBadRequest)
		case 2:
			// first param is the name of the invalid parameter; second param is more detail on why it's invalid
			jsonMsgResponse(w, r, "0003", fmt.Sprintf("invalid parameter: %s; %s", invalidParam[0], invalidParam[1]), http.StatusBadRequest)
		default:
			// use the first param only
			jsonMsgResponse(w, r, "0003", fmt.Sprintf("invalid parameter: %s", invalidParam[0]), http.StatusBadRequest)
		}
	case api.ErrBadRequest:
		jsonMsgResponse(w, r, "0004", "bad request", http.StatusBadRequest)
	default:
		jsonMsgResponse(w, r, "0001", "internal server error", http.StatusInternalServerError)
	}
}
