package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApp() (*Application, error) {
	logger := log.New(os.Stdout, "From logger ", log.Ldate|log.Ltime)

	return &Application{
		Logger: logger,
	}, nil
}

