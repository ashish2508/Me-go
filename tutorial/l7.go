package main

import "fmt"

func main() {
	fname, _ := getNames()

	fmt.Println("Welcome to GOlang,", fname)

}

func getNames() (string, string) {
	return "John", "Doe"
}
