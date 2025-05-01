// S T R U C T

package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {

	return true
}

func test(mToSend messageToSend) {

}

func main() {

}
