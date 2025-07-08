package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webfont video.")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:4000/"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("The status code is:", response.StatusCode)

	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:4000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	{
	"courseName":"Let's go with golang.",
	"price": 0,
	"platform": "learncode"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:4000/postform"

	// formData

	data := url.Values{}

	data.Add("firstname", "Dhirya")
	data.Add("lastname", "Pandya")
	data.Add("email", "dhairya@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
