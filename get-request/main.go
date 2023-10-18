package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling Get Request ......")
	PerformGetRequest()

}

func PerformGetRequest() {
	const url = "https://dummyjson.com/products/1"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecontent, _ := responseString.Write(content)

	fmt.Println(bytecontent)
	fmt.Println(responseString.String())

}
