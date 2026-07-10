package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/store"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(app app.Application, w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username or password can not be empty.", http.StatusBadRequest)
		return
	}

	if err := app.DB.User.CreateUser(user.Username, user.Password); errors.Is(store.UserDuplicateError, err) {
		http.Error(w, fmt.Sprintf("already a user with username of %s exist.", user.Username), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
