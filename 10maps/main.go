package main
import "fmt"


func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "RubyonRails"
	languages["PH"] = "Python"


	fmt.Println(languages) 


	delete(languages, "JS")
	fmt.Println("JS shorts for ", languages["PH"])
	fmt.Println(languages)



	// loops are interesting

	for key, value := range languages {
		fmt.Printf("For key %v, value is is %v \n", key, value)
	}


}