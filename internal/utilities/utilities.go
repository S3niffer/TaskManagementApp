package utilities

import (
	"encoding/json"
	"net/http"
	"os"
)

func JsonMessage(m string, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", ": application/json")

	return json.NewEncoder(w).Encode(m)
}

func IsFileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return false
}
