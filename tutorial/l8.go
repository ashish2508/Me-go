// STRUCT

package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending messages : '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("=======================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for Signing up",
	})

	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard",
	})

	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})

}
