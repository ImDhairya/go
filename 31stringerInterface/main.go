package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// func () String() string {

// }

// this is the stringer interface something that return the strign in the value and is named String() in the name

func (p Person) String() string {
	return fmt.Sprintf("%v  (%v years ) ", p.Name, p.Age)
}

func main() {
	p := Person{"Bobby", 22}

	q := Person{"Ajitesh", 33}

	fmt.Println(p, q)
}
