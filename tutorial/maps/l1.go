//make keyword together with map

package main

import "fmt"

func main(){
	ages:= make(map[string]int)
	ages["john"]=38
	ages["mak"]=33
	ages["jellu"]=34
	ages["jonny"]=35
	ages["johattt"]=36
	
	for name, age:=range ages{
		fmt.Printf("Names is %v and age is %v\n",name, age)
	}
}
