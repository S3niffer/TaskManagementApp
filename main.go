package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/router"
)

func main() {
	var err error

	application, err := app.New()

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
