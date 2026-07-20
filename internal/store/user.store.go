package store

import (
	"database/sql"

	"github.com/s3niffer/taskmanagementapp/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) UserStore {
	return UserStore{
		db: db,
	}
}

func (us UserStore) AddUser(username, email, password string, u *models.User) error {
	query := `
	INSERT INTO users (username,email,password_hash)
	VALUES ($1,$2,$3)
	RETURNING id,created_at,updated_at;
	`

	err := us.db.QueryRow(query, username, email, password).Scan(&u.ID, &u.Create_at, &u.Updated_at)
	if err != nil {
		return err
	}

	return nil
}
