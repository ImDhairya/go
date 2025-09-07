package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Not correct method", http.StatusNotAcceptable)
		return
	}

	fmt.Fprintln(w, "hello ")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	fmt.Fprintf(w, "Post request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "The name is %s \n", name)
	fmt.Fprintf(w, "The address is %s \n ", address)

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandle)

	fmt.Printf("Starting the server at 8080 \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
