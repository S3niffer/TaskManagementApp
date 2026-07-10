package store

type Store struct {
	User UserStore
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
