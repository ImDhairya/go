package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang ")

	presentTime := time.Now()

	fmt.Println("present Time is" , presentTime.Format("01-02-2006 15:04:05 Monday"))


	createdData := time.Date(2020, time.January, 10, 10, 32,0,0,time.UTC)

	fmt.Println("Created date", createdData.Format("01-02-2006 Monday" ))
}