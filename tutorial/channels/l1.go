package main

import (
	"fmt"
	"time"
)

func main(){
	 // Create a channel
 ch := make(chan string)

 // Start a goroutine to send data to the channel
 go func() {
	for i := range 5 {
	 ch <- fmt.Sprintf("Message %d", i)
	 time.Sleep(1 * time.Second) // Simulate some work
	}
	close(ch) // Close the channel when done
 }()

 // Receive data from the channel
 for msg := range ch {
	fmt.Println("Received:", msg)
 }
fmt.Println("Channel closed, exiting.")
}
