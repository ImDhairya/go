package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {

	welcome := "Welcome to user input program."
	fmt.Printf("This is %v message\n", welcome)

	// pointer

	reader := bufio.NewReader(os.Stdin)

	// S and N are capital because of the public variables they are 

	fmt.Println("Enter the rating for our pizza: ")

	// fmt.Printf("", var)
	// treat the errors as variables, so there is no try catch


	// comma ok or comma error
	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	fmt.Printf("Type of this rating is %T\n", input)


}