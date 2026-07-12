package main

import (
	"log"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
)

func main() {
	application := app.New()

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	http.HandleFunc("/health", application.HealthCheck)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
