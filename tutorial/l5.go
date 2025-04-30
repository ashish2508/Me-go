package main

import "fmt"

func concat(s1, s2 interface{}) interface{} {
	switch x := s1.(type) {
	case int:
		switch y := s2.(type) {
		case int:
			return x + y
		case string:
			return fmt.Sprintf("%d%s", x, y)
		}
	case string:
		switch y := s2.(type) {
		case int:
			return fmt.Sprintf("%s%d", x, y)
		case string:
			return x + y
		}
	}
	return nil

}

func main() {
	fmt.Println(concat("hi", " ashish"))
	fmt.Println(concat("how ", "are you"))
	fmt.Println(concat("i am good", " how are you"))
	fmt.Println(concat(2, 3))
	fmt.Println(concat(3, "ashish"))
	fmt.Println(concat("ashish", 99))
}
