package main

import "fmt"

func check() {
	defer fmt.Println("world")
	fmt.Println("Hello.")
}

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")


	fmt.Println("Hello.")
	myDefer()
}

func myDefer() {
	// for i := 0 ; i < 5 ; i++ {
	// 	defer fmt.Println(i)
	// }

	for i := range[5]int {} {
		defer fmt.Println(i)
	}
}


// hello
// 