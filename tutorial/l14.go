//custom error
package main 
import "fmt"

type useError struct {
	name string
}

func (e useError)  Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

func main(){
	e:= useError{"Ashish"}
	fmt.Printf("%v\n",e)
}
