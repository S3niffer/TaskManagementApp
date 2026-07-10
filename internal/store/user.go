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

func (store *UserStore) CreateUser(u string, p string) (int, error) {
	store.db.Lock()
	defer store.db.Unlock()
	for _, v := range store.db.Users {
		if v.Username == u {
			return 0, UserDuplicateError
		}
	}

	store.db.Users = append(store.db.Users, user{Username: u, Password: p, ID: store.db.ModifiedTimes + 1})
	store.db.ModifiedTimes++

	err := store.db.SaveToFile()
	if err != nil {
		return 0, err
	}

	return store.db.ModifiedTimes, nil
}

func (store *UserStore) GetAllUsers() []user {
	store.db.RLock()
	defer store.db.RUnlock()

	users := make([]user, len(store.db.Users))
	copy(users, store.db.Users)

	return users
}
