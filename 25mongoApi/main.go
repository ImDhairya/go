package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ImDhairya/go/router"
)

func main() {
	fmt.Println("Mongodb API")

	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4040", r))

	fmt.Println("Listening at port 4000")

}
