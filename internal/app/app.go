package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/s3niffer/taskmanagementapp/internal/api"
	"github.com/s3niffer/taskmanagementapp/internal/store"
	"github.com/s3niffer/taskmanagementapp/migrations"
)

type Application struct {
	DB      *sql.DB
	UserApi api.UserApi
}

func New() (Application, error) {
	db, err := store.ConnectToDB()
	if err != nil {
		return Application{}, err
	}

	err = store.Migrate(db, "postgres", ".", migrations.FS)
	if err != nil {
		return Application{}, err
	}

	return Application{
		DB:      db,
		UserApi: api.NewUserApi(store.NewUserStore(db)),
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
