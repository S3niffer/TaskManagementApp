package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/s3niffer/taskmanagementapp/internal/models"
	"github.com/s3niffer/taskmanagementapp/internal/store"
)

type UserApi struct {
	Store store.UserStore
}

func NewUserApi(store store.UserStore) UserApi {
	return UserApi{
		Store: store,
	}
}

func (u UserApi) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Error RegisterUser: %s\n", err.Error())
		http.Error(w, fmt.Errorf("parsing the body: (%w)", err).Error(), http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Username == "" || user.Password_hash == "" {
		fmt.Print("Error RegisterUser: empty values")
		http.Error(w, "none of email,username,password can be empty.", http.StatusBadRequest)
		return
	}

	err = u.Store.AddUser(user.Username, user.Email, user.Password_hash, &user)
	if err != nil {
		fmt.Printf("Error RegisterUser: add to db %s", err.Error())
		http.Error(w, "couldn't insert into db.", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(models.NewUser{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Create_at:  user.Create_at,
		Updated_at: user.Updated_at,
	})
}

func (u UserApi) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Error login user: %s\n", err.Error())
		http.Error(w, fmt.Errorf("parsing the body: (%w)", err).Error(), http.StatusBadRequest)
		return
	}

	_, pass, err := u.Store.FindUser(user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "No such user has been found.", http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Print("ERROR", err)
		return
	}

	if pass == user.Password_hash {
		w.Write([]byte("youre logged in."))
		return
	}
	w.Write([]byte("youre not logged in."))
}
