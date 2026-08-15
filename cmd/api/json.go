package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

// Go data structure to JSON
func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// read JSON from user and decode into the given data  structure - data
func ReadJSON(w http.ResponseWriter, status int, r *http.Request, data any) error {
	maxBytes := 1_048_578                                    // 1 mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // to prevent security attacks that comes with large-size data.
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return WriteJSON(w, status, &envelope{Error: message})
}
