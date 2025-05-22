//formatting string 

package main 
 import ("fmt")

func msg (name string, age int8) string {
	s:= fmt.Sprintf("%v is %v years old.", name, age)
	return s 
}
func msg2 (name string, age int8) string {
	s:= fmt.Printf("%v is %v years old.", name, age)
	return s 
}
func main (){
	fmt.Printf("%s",msg("ashish", 20))
	//won't work
	fmt.Printf("%s",msg2("ash", 20))
}
