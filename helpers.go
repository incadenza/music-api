package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type response map[string]interface{}

func writeJSON(w http.ResponseWriter, status int, data response) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func successResponse(w http.ResponseWriter, r *http.Request, data response) {
	writeJSON(w, 200, data)
}

func parseIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return id, nil
}
