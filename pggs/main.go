package main

import (
	"fmt"
	// "golang.org/x/text/number"
	// "time"
)

// func main() {
// 	var i int
// 	var f float32
// 	var b bool
// 	var s string

// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// func main() {
// 	t := time.Now()

// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good Morning!!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good Afternoon")
// 	default:
// 		fmt.Println("Good Evening!!!")
// 	}
// }

// func main() {
// 	sum := 10
// 	fmt.Printf("\n Hello \n World ")
// 	defer fmt.Printf("\n Okayy \n ")
// 	defer fmt.Printf("\n Will be at the end of the code \n")

// 	fmt.Printf("Hii %v %T", sum, sum)
// }

type Vertex struct {
	X int
	Y int
}

type SecondStruct struct {
	X, Y int
}

var (
	v1 = Vertex{1, 44}
	v2 = Vertex{X: 15}

	v3 = Vertex{}
	p  = &Vertex{22, 56}
)

const (
	val = "world "
)

// func main() {

// 	var (
// 		Gf = "ehhehehe"
// 	)
// 	// v := Vertex{34, 3}
// 	// myStruct := SecondStruct{43, 67}

// 	// good := &myStruct

// 	// good.X = 78
// 	// fmt.Println(good)
// 	fmt.Println(p, "The pointer ", v1, v2, v3, val, Gf)

// }

// func main() {
// 	s := []int{2, 3, 5, 7, 9}
// 	a := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(s, a)
// 	printSlices(s)
// }

// func printSlices(s []int) {
// 	fmt.Printf("len=%d cap %d %v\n", len(s), cap(s), s)
// }

// func main() {
// 	var s []int

// 	fmt.Println(s, len(s), cap(s))
// 	// the value of the slice will be nill
// 	if s == nil {
// 		fmt.Println("nil!")
// 	}
// }

type VarStruct struct {
	name     string
	phone    int
	address  string
	marks    []int
	subjects [3]string
}

// func main() {

// 	person := VarStruct{
// 		"Dhairya",
// 		3323,
// 		"fefeffef",
// 		[]int{23, 2, 12},
// 		[3]string{"hind ", "env", "fefef"},
// 	}

// 	s := []int{}
// 	fmt.Println(person, "Heyya")
// 	fmt.Println("The zero value of the slice is nill ", s)
// 	fmt.Printf("\n %v a slice has zero value as nil %v \n ", s, s)

// }

// func main() {
// 	a := make([]int, 5)

// 	b := make([]string, 5)

// 	b[0] = "hello worl "

// 	b[2] = "This is a string."

// 	fmt.Println(a, b, "The length and capacity by using make in go.")

// }

// func main() {

// 	board := [][]string{
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 	}

// 	board[0][0] = "X"

// 	fmt.Printf(" %v \n", board )

// }

// func main() {
// 	var nums []int

// 	for i := 0; i < 20; i++ {
// 		nums = append(nums, i)
// 		fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums))
// 	}

// }

func main() {

	// var s []int
	// printSlices(s)
	// s = append(s, 99, 32)
	// printSlices(s)
	// s = append(s, 55)
	// printSlices(s)

	// for i := 1; i < 10; i++ {
	// 	s = append(s, i)
	// 	printSlices(s)
	// }

	var betterSLice = make([]int, 0, 10)

	for i := 1; i < 10; i++ {
		betterSLice = append(betterSLice, i)
		printSlices(betterSLice)
	}

}

func printSlices(s []int) {
	fmt.Printf("Capacity %d Length %d, %v \n ", cap(s), len(s), s)
}
