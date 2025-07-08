package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://strawbrewery.site/"

func main() {
	fmt.Println("Strawbrewery web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close() // callers responsibility to close the connection.

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)


	fmt.Println(content)
}
