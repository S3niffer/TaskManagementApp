package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/router"
)

func main() {
	var err error

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 20,
		Handler:      router.New(),
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Println("couldn't run the server. :(")
	}
}
