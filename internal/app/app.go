package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/s3niffer/taskmanagementapp/internal/api"
	store "github.com/s3niffer/taskmanagementapp/internal/repository"
	"github.com/s3niffer/taskmanagementapp/migrations"
)

type Application struct {
	Logger *log.Logger
	User   *api.UsersHandler
	DB     *sql.DB
}

func NewApp() (*Application, error) {
	logger := log.New(os.Stdout, "From logger ", log.Ldate|log.Ltime)

	pgDB, err := store.Open(logger)
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		logger.Panic(err)
	}

	userStore := store.NewPostgresUserStore(pgDB)

	userHandler := api.NewUsersHandler(userStore)

	return &Application{
		Logger: logger,
		User:   userHandler,
		DB:     pgDB,
	}, nil
}

func (Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available.\n")
}
