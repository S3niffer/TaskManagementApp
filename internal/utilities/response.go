package utilities

import (
	"encoding/json"
	"net/http"
)

func JsonMessage(m string, w http.ResponseWriter) {
	w.Header().Add("Content-Type", ": application/json")

	if err := json.NewEncoder(w).Encode(m); err != nil {
		http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	}
}
