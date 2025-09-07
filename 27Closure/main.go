package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


// closure means the inner function will have the access of the outer functions variables and that too even after the outher function (returned )has finished executing
