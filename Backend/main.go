package main

import (
	"github.com/ashish2508/Me-go/Backend/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Works")
	app, err := app.NewApplication()
	if err != nil{
		panic(err)
	}
	
	
}
