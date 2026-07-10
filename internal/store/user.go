package store

type UserStore struct {
	db *DataBase
}

func (store *UserStore) CreateUser(u string, p string) error {

	store.db.users = append(store.db.users, user{username: u, password: p, id: store.db.modifiedTimes + 1})
	store.db.modifiedTimes++

	return nil
}
