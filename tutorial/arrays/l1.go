package main
import "fmt"

/*
in go arrays have fixed[static] size 
*/
func main(){
	primes:=[6]int{2,3,5,7,11,13}
	for i:=1;i < len(primes)+1;i++{
		fmt.Printf("Primes %d is %d\n",i,primes[i-1])
	}
}
