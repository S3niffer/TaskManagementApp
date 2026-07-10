package store

import (
	"errors"
)

type UserStore struct {
	db *DataBase
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

var UserDuplicateError = errors.New("Duplicate username.")

func (store *UserStore) CreateUser(u string, p string) error {

	for _, v := range store.db.Users {
		if v.Username == u {
			return UserDuplicateError
		}
	}

	store.db.Users = append(store.db.Users, user{Username: u, Password: p, ID: store.db.ModifiedTimes + 1})
	store.db.ModifiedTimes++

	err := store.db.SaveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (store *UserStore) GetAllUsers() []user {
	return store.db.Users
}
