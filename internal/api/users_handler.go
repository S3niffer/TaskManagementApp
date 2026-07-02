package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	store "github.com/s3niffer/taskmanagementapp/internal/repository"
)

type UsersHandler struct {
	userStore store.UserStore
}

func NewUsersHandler(userStore store.UserStore) *UsersHandler {
	return &UsersHandler{
		userStore: userStore,
	}
}

func (user *UsersHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	paramsId := chi.URLParam(r, "id")
	if paramsId == "" {
		http.Error(w, "User id is not a found.", http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		http.Error(w, "User id is not a number.", http.StatusBadRequest)
		return
	}

	userStruct := &store.User{ID: int(userId)}

	if err = user.userStore.FindUserById(userStruct); err != nil {
		fmt.Println(err)
		http.Error(w, "Couldn't find user with provided id.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(userStruct); err != nil {
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}
}

func (u *UsersHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user store.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	createdUser, err := u.userStore.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}
