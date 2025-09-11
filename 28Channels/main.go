package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang ")

	myCh := make(chan int) 
	wg := &sync.WaitGroup{}

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5

		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

	// myCh <- 5
	// fmt.Println(<-myCh)

}
