package main

import "fmt"

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \b \n", v, v*v)
	case string:
		fmt.Printf("%q is %v bytes long \n ", v, len(v))
	default:
		fmt.Println("I don't know about the type.")
	}
}

func main() {

	do(21)
	do("hello")
	do(true)
}
