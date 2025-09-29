package main

import (
	"fmt"
	"net/http"
	"time"
	"flag"
	"github.com/ashish2508/Me-go/Backend/internal/app"
)

func main() {
	fmt.Println("Works")
	
	var port int 
	flag.IntVar(&port, "port", 8080,"go backend server port")
	flag.Parse()
	
	app, err := app.NewApplication()
	if err != nil{
		panic(err)
	}
	
	
	
	http.HandleFunc("/health", HealthCheck)   
	server := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		IdleTimeout: time.Minute,
		ReadTimeout: time.Second * 10,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("We are running on Port: %d\n",port)
	
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Status is available")
}
