package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/handler"
)

func main() {
	var port = 8080
	flag.IntVar(&port, "port", 8080, "Server app port.")
	flag.Parse()

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer application.DB.Close()

	handler := handler.New(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	fmt.Printf("\nServer in up on port: %d\n", port)

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
