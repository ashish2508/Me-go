package main

import (
	"fmt"
	"github.com/ashish2508/mystrings"
)

func main(){
	text := "Hello world"

	fmt.Println(text)
	fmt.Println(mystrings.Reverse(&text))
}
