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
	RETURNING id,created_at;
	`

	err := us.db.QueryRow(query, username, email, password).Scan(&u.ID, &u.Create_at)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStore) FindUserByUsername(username string) (int, string, error) {
	var id int
	var pass string

	query := `
		SELECT id,password_hash FROM users WHERE username = $1;
	`

	err := us.db.QueryRow(query, username).Scan(&id, &pass)
	if err != nil {
		return 0, "", err
	}

	return id, pass, err

}

func (us UserStore) FindUserById(id int) (models.NewUser, error) {
	var user models.NewUser
	query := `
		SELECT * FROM users WHERE id = $1;
	`

	err := us.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password_hash, &user.Create_at)
	if err != nil {
		return user, err
	}

	return user, err

}
