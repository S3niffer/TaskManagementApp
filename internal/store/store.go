package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // <-- Required
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	User UserStore
}

func New() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=1 dbname=task_managment sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database...")
	return db, nil
}
