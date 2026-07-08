package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	if json.NewEncoder(w).Encode(user) != nil {
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
}
