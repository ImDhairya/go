package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang.")

	// no inheritance, parents , no super

	dhairya := User{"Dhairya", "abc@gmail.com", false, 22}


	fmt.Println(dhairya)

	fmt.Printf("Dhairya details are: %+v\n", dhairya)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
