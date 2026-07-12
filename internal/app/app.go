package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Application struct {
	DB *sql.DB
}

func New() (Application, error) {
	db, err := connectToDB()
	if err != nil {
		return Application{}, err
	}

	return Application{
		DB: db,
	}, nil
}

func (Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}

func connectToDB() (*sql.DB, error) {
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
