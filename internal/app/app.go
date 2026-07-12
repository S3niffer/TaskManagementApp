package app

import (
	"encoding/json"
	"net/http"
)

type Application struct {
}

func New() Application {
	return Application{}
}

func (Application) HealthCheck(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
