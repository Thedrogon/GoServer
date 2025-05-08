package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application,error){
	logger:= log.New(os.Stdout,"",log.Ldate|log.Ltime)  //here "" says about logging semnatics where in this case its gonna follow default semantics

	app := &Application{ //modifying our actual application
		Logger: logger,
	}

	return app,nil
}