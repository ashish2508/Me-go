package main

import "fmt"

func main(){
	fruits:= [][]string{
		{"apple","red"},
		{"mango","orangish-yellow"},
		{"banana","yellow"},
		{"melon","greenish-yellow"},
		{"muskmelon","green"},
		{"watermelon","dark-green"},
		{"lemon","yellow"},
}
		
	for i,fruit :=  range fruits {
			fmt.Printf("%d. %v is a Fruit and its color is %v\n",i+1,fruit[0],fruit[1])
		}
}
