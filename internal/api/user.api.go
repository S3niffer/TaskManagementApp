package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/s3niffer/taskmanagementapp/internal/models"
	"github.com/s3niffer/taskmanagementapp/internal/store"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password_hash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "password hash", http.StatusInternalServerError)
	}

	err = u.Store.AddUser(user.Username, user.Email, string(hash), &user)
	if err != nil {
		fmt.Printf("Error RegisterUser: add to db %s", err.Error())
		http.Error(w, "couldn't insert into db.", http.StatusInternalServerError)
		return
	}

	token, err := getJwtToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(models.NewUser{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Create_at:   user.Create_at,
		Updated_at:  user.Updated_at,
		AccessToken: token,
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

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password_hash))
	if err != nil {
		http.Error(w, "Password is wrong", http.StatusUnauthorized)
		return
	}

	token, err := getJwtToken(user.ID)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	})
}

func getJwtToken(id int) (string, error) {
	claims := jwt.MapClaims{
		"userID": id,
		"exp":    time.Now().Add(time.Minute * 2).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return token, nil
}
