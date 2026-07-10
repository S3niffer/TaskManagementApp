package store

import (
	"errors"
)

type UserStore struct {
	db *DataBase
}

type user struct {
	username string
	password string
	id       int
}

var UserDuplicateError = errors.New("Duplicate username.")

func (store *UserStore) CreateUser(u string, p string) error {

	for _, v := range store.db.users {
		if v.username == u {
			return UserDuplicateError
		}
	}

	store.db.users = append(store.db.users, user{username: u, password: p, id: store.db.modifiedTimes + 1})
	store.db.modifiedTimes++

	return nil
}
