package main 
import (
"fmt"
"time"
)
func main(){
 numSentCh := make(chan int,10)
 done := make(chan bool)

total := 0
go func() {
	for {
		numSent,ok := <-numSentCh
	if !ok {
		break
	}
	total += numSent
	time.Sleep(1 * time.Microsecond) // Simulate some work
	fmt.Println("Received:", numSent)
	}
	done <- true
	}()
	numSentCh <- 5
	numSentCh <- 3
	numSentCh <- 7
	numSentCh <- 20
	numSentCh <- 21
	numSentCh <- 2
	numSentCh <- 25
	numSentCh <- 23
	numSentCh <- 22
	numSentCh <- 12
	
	close(numSentCh)
	<-done
	fmt.Println("Total messages sent:", total)
	fmt.Println("Processing complete!")
}
