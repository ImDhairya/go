package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in go lang")

	dhairya := User{"Dhairya", "dhairya@gmail.com", true, 22}

	fmt.Println(dhairya, "The name details.")

	dhairya.GetStatus()

}


func (u User) GetStatus() {
fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("Email of this user is: ", u.Email)
}