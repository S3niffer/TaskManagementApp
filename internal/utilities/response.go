package utilities

import (
	"encoding/json"
	"net/http"
)

func JsonMessage(m string, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", ": application/json")

	return json.NewEncoder(w).Encode(m)
}
