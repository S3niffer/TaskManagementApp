package store

import (
	"database/sql"
	"errors"
	"log"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

var UserDuplicateError = errors.New("Duplicate username.")

func (store *UserStore) CreateUser(u string, p string) (int, error) {
	query := `
	INSERT INTO users(username,password_hash,email)
	VALUES($1,$2,$3)
	RETURNING id
	`
	id := 0

	err := store.db.QueryRow(query, u, p, "test@Gmnail2.com").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (store *UserStore) GetAllUsers() []User {
	query := `
	SELECT * FROM users
	`

	rows, err := store.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// for rows.Next() {
	// 	var u User

	// 	err := rows.Scan(&u.ID,)
	// }

	return []User{}
}
