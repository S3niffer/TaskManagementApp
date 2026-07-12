package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/s3niffer/taskmanagementapp/internal/handler"
	"github.com/s3niffer/taskmanagementapp/internal/store"
	"github.com/s3niffer/taskmanagementapp/internal/utilities"
)

type Application struct {
	DB      store.Store
	Handler handler.Handler
}

func New() (Application, error) {
	db, err := store.New()
	if err != nil {
		fmt.Println(err)
		fmt.Println("couldn't create the database. :(")
		os.Exit(1)
	}

	handler := handler.New(db)

	return Application{
		DB:      db,
		Handler: handler,
	}, nil
}

func (Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", ": application/json")

	// if err := json.NewEncoder(w).Encode("It looks fine. :)"); err != nil {
	// 	http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	// }

	if err := utilities.JsonMessage("It looks fine. :)", w); err != nil {
		http.Error(w, "Something went wrong!.", http.StatusInternalServerError)
	}
}
