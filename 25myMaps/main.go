package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// fmt.Println("This is for maps ")

	// m = make(map[string]Vertex)

	// m["Bell Lamp"] = Vertex{
	// 	40.3324343, 434.3788,
	// }

	// fmt.Println(m["Bell Lamp"])

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value ", m["Answer"])

	m["Mark"] = 43

	for key, value := range m {
		fmt.Println(key, value)
	}

	// for _, value := range m {
	// 	fmt.Println(value, "Fefe")
	// }

	// for key, _ := range m {
	// 	fmt.Println(key, "efef")
	// }

	fmt.Println("The map is ", m)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// This is called as the value ok syntaxt.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// there are only 2 ways of declaring a map

	// 1) Make
	// m := make(map[string]int)
	// m[Answer] = 55

	// 2) Literal syntaxt
	// var m = map[string]Vertex{
	// 	"Bell Labs ": Vertex{
	// 		32.32, 90.00,
	// 	},
	// 	"Google": Vertex{
	// 		67.21, 44.00,
	// 	},
	// }

	// another example

	// var m = map[string]Vertex {}

	// you can not declare a map with var syntaxt that we do in other data types
	// wrong
	//  var m map[string]int
	// m["Key1"] = 67
	// this cant be done
}
