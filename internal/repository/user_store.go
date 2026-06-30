package store

import "database/sql"

type User struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_hash string `json:"password_hash"`
	Created_at    string `json:"create_at"`
	Updated_at    string `json:"updated_at"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db}
}

type UserStore interface {
	CreateUser(*User) (*User, error)
}

func (pg PostgresUserStore) CreateUser(user *User) (*User, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
	INSERT INTO users (username,email,password_hash)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err = tx.QueryRow(query, user.Username, user.Email, user.Password_hash).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return user, nil
}
