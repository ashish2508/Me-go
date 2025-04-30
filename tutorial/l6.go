package main

import "fmt"

func main() {
	sendSoFar := 430
	const sendsToAdd = 25
	sendSoFar = increm(sendSoFar, sendsToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")
}

func increm(sendSoFar, sendsToAdd int) int {
	sendSoFar += sendsToAdd
	return sendSoFar
}
