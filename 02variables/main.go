package main

import "fmt"

const password int = 3289
// small p means password i local variable
// capital P means the value is a global variable	 
func main() {
	var username string = "Dhairya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true

	fmt.Println(isLoggedIn)
	// fmt.Printf("This is the type of boolean %T \n ", isLoggedIn)
	fmt.Printf("This is the logged in value that is there. %v \n", isLoggedIn)


	var smallInt int = 3448232

	fmt.Printf("This is where we are goonna learn about the small numbers %v\n", smallInt)


	
	// Default values and some alliaces herefl aewffeawfewfewf
	var anotherVariable string
	fmt.Printf("The password %v \n", password)
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is %T \n", anotherVariable)


}