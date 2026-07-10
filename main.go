package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/router"
	"github.com/s3niffer/taskmanagementapp/internal/store"
)

func main() {
	var err error

	db, err := store.New()
	if err != nil {
		fmt.Println("couldn't create the database. :(")
		os.Exit(1)
	}

	application, err := app.New(db)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 20,
		Handler:      router.New(application),
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err)
		fmt.Println("couldn't run the server. :(")
	}
}
