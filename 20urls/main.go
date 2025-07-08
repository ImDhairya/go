package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fffs83a33f3#$@@1"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myUrl, "My url is here")

	// parsing

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("The scheme of this result is: ", result.Scheme)

	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params is. %T\n", qparams)

	fmt.Println("The course name is: ", qparams["coursename"])
	fmt.Println("The payment id is: ", qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println(val, "HELLOWWW")
	}
}
