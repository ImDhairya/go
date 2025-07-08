package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is all about slices")

	var fruitLists = [] string{"banana", "apple", "mango"}

	fmt.Printf("This is a list of fruits. %s\n", fruitLists)

	// fruitLists = append(fruitLists, "banana", "chickoo")

	fruitLists = append(fruitLists[1:])

	// fmt.Println(fruitLists)


	highScors := make([]int,4)

	highScors[0] = 234
	highScors[1] = 134
	highScors[2] = 534
	highScors[3] = 34

	highScors = append(highScors, 32,21,3)
	// fmt.Println(highScors)

	sort.Ints(highScors)
	// fmt.Println(highScors)
	// fmt.Println(sort.IntsAreSorted(highScors))


	// remove data from slice based on index value

	var courses = []string{"React", "Swift", "Python"}
	fmt.Println(courses)

	index := 1

	courses = append(courses[:index], courses[index + 1:]...)

	fmt.Println(courses)



}