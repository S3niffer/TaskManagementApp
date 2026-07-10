package store

type Store struct {
	User UserStore
}

type user struct {
	username string
	password string
	id       int
}

type DataBase struct {
	users         []user
	modifiedTimes int
}

func New() (Store, error) {
	db := &DataBase{}

	return Store{
		User: UserStore{db},
	}, nil
}
