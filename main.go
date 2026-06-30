package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/handler"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Used for backend server port")
	flag.Parse()

	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	defer application.DB.Close()

	handler := handler.SetUpHandler(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      handler,
	}

	application.Logger.Printf("We are running at port: %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatal(err)
	}

}
