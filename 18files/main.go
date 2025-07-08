// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	fmt.Println("welcome to files in go lang.")

// 	content := "This needs to go in file - lco."

// 	file, err := os.Create("./myDhairya.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	length, err := io.WriteString(file, content)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("The length is : ", length)
// 	defer file.Close()

// 	readFile("./myDhairya.txt")

// }

// func readFile(fileName string) {

// 	dataByte , err :=  ioutil.ReadFile(fileName)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Text data inside the file is \n ", dataByte)

// }

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const content = "This is the content that is supposed to go in the file."

	file, err := os.Create("./dhairya.txt")

	if err != nil {
		panic(err)

	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length, "This is the length of the file.")

	defer file.Close()

	readFile("./dhairya.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The databyte size is here. %v \n", databyte)
	fmt.Printf("This is in string format. %s \n", string(databyte))
}
