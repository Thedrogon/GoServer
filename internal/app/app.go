package app

import (
	"fmt"
	"log"
	"net/http"
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

//a health check function - its a http handler so it has a signature

func (app *Application)HealthCheck(w http.ResponseWriter, r *http.Request) { //here r is a pointer but w isnt because later we will understand that client data i.e request has alot more data to process specially during middleware use.
	fmt.Fprintf(w , " Status is available ")
}