package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json data Handle ..... ")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []Course{
		{"GoCourse", 40, "Local", "abc123", []string{"web-dev", "go"}},
		{"ReactCourse", 40, "Local", "abc123", []string{"web-dev", "react"}},
		{"PHPCourse", 40, "Local", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(` 
	{
		"coursename": "ReactCourse",
		"Price": 40,
		"website": "Local",
		"tags": [
					"web-dev",
					"react"
		]
	}
`)
	var course Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid Json")

		json.Unmarshal(jsonDataFromWeb, &course)

		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("Not valid json")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of value is %T\n", key, val, val)
	}
}
