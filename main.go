package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Thedrogon/v1/Goserver/internal/app"
)

func main() {

	var port int
	flag.IntVar(&port,"port",8080,"Go backend port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	
	http.HandleFunc("/health", HealthCheck) //in Go functions are first class citizens that is we can use them as variables or passing them as params in another method.
	
	server := &http.Server{ //a struct that defines the http properties
		Addr:         fmt.Sprintf(":%d",port),
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

//a health check function - its a http handler so it has a signature
func HealthCheck(w http.ResponseWriter, r *http.Request) { //here r is a pointer but w isnt because later we will understand that client data i.e request has alot more data to process specially during middleware use.
	fmt.Fprintf(w , " Status is available ")
}
