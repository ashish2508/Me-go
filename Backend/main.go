package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ashish2508/Me-go/Backend/internal/app"
)

func main() {
	fmt.Println("Works")
	app, err := app.NewApplication()
	if err != nil{
		panic(err)
	}
	
	app.Logger.Println("We are running out app")
	
	server := &http.Server{
		Addr: ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: time.Second * 10,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
	
	
}
