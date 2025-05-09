package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Thedrogon/v1/Goserver/internal/app"
	"github.com/Thedrogon/v1/Goserver/internal/routes"
)

func main() {

	var port int
	flag.IntVar(&port,"port",8080,"Go backend port") //here default value is 8080 when -port verbose is not used other wise takes in the port value from terminal when starting the server ==> "go run main.go -port 8082"
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	
	//http.HandleFunc("/health", app.HealthCheck) //in Go functions are first class citizens that is we can use them as variables or passing them as params in another method.

	r := routes.SetRoutes(app)
	
	server := &http.Server{ //a struct that defines the http properties
		Addr:         fmt.Sprintf(":%d",port),
		Handler:				r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	app.Logger.Printf("We are running our app on port %d\n",port)


	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}


