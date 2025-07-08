package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Welcome to JSON video.")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this as json

	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	  {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "Platform": "LearnCodeOnline.in",
                "Password": "abc123",
                "Tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var dataCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("This is a valid json format data.")
		json.Unmarshal(jsonDataFromWeb, &dataCourse)

		fmt.Printf("%#v\n", dataCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID.")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
}
