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

type newUser struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func CreateUser(app app.Application, w http.ResponseWriter, r *http.Request) {
	var user User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&user); err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Bad request: %s)", err.Error()), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username or password can not be empty.", http.StatusBadRequest)
		return
	}

	id, err := app.DB.User.CreateUser(user.Username, user.Password)
	if errors.Is(store.UserDuplicateError, err) {
		http.Error(w, fmt.Sprintf("already a user with username of %s exist.", user.Username), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := newUser{Username: user.Username, ID: id}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, fmt.Sprintf("Internal Error: (%s)", err.Error()), http.StatusInternalServerError)
	}
}

func GetAllUsers(app app.Application, w http.ResponseWriter) {
	users := app.DB.User.GetAllUsers()

	w.Header().Add("Content-Type", ": application/json")

	json.NewEncoder(w).Encode(users)
}
