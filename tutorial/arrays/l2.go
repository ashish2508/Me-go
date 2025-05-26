
package main

import "fmt"

func main() {	

	//slices
	a1 := []int{1, 2, 3, 4, 5}
	b :=a1[0:3:1] // slice from index 0 to 2
	a:=fmt.Sprintf("a1:%d", a1)
	fmt.Println(a)
	fmt.Println("b:", b)
}
