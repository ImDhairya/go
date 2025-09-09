package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup

var mut sync.Mutex // pointer

func main() {

	// greeter("Hello")
	// go greeter("World")

	website := []string{
		"https://google.com",
		"http://pargs.org/",
		"https://amazon.com",
		"https://nike.com",
		"https://github.com",
		"https://instagram.com",
	}

	for _, item := range website {
		go getStatusCode(item)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Opps in endpoint")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
	defer wg.Done()

}
