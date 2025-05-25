package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost:=0.0
	for i:=0.0;i<float64(numMessages);i++ {
			totalCost+=1.0+(0.01*i)
	}
	return totalCost
}

func main(){
	var numMessages int
	fmt.Print("Enter your number of Messages: ")
	fmt.Scanln(&numMessages)
	cost:=bulkSend(numMessages)
	fmt.Printf("Total cost is : %.2f",cost);
}
