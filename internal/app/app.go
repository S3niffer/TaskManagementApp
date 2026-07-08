package app

import (
	"encoding/json"
	"net/http"
)

type Application struct {
}

func New() (*Application, error) {
	return &Application{}, nil
}

func (Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", ": application/json")

	if err := json.NewEncoder(w).Encode("It looks fine. :)"); err != nil {
		http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	}
}
