package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/router"
)

func main() {
	var err error

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer application.DB.Close()

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
