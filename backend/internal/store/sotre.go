package store

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5433 user=postgres password=mypassword dbname=task_management sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Open db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Open db: %w", err)
	}

	fmt.Print("Database is up.")

	return db, nil
}

func Migrate(db *sql.DB, driver, dir string, fs embed.FS) error {
	err := goose.SetDialect(driver)
	if err != nil {
		return fmt.Errorf("Migrate (SetDialect) error: %w", err)
	}

	goose.SetBaseFS(fs)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	fmt.Println("Running migrations ...")
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("Migrate (Up) error: %w", err)
	}
	fmt.Println("migrations ended.")

	return nil
}
