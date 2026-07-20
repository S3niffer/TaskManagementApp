package api

import (
	"encoding/json"
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
