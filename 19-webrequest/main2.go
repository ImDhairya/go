package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main2() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://example.com/api"
	method := "POST"
	payload := strings.NewReader(`{"key1" : "value1" , "key2" : "value2"}`)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("Error creating request", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer YOUR_TOKEN")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error making request.", err)

		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Println(string(body))
}
