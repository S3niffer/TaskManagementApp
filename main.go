package main

import (
	"log"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/handler"
)

func main() {
	application := app.New()

	handler := handler.New(application)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
