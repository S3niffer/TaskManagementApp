package app

import (
	"net/http"

	"github.com/s3niffer/taskmanagementapp/internal/utilities"
)

type Application struct {
}

func New() (*Application, error) {
	return &Application{}, nil
}

func (Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", ": application/json")

	// if err := json.NewEncoder(w).Encode("It looks fine. :)"); err != nil {
	// 	http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	// }

	if err := utilities.JsonMessage("It looks fine. :)", w); err != nil {
		http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	}
}
