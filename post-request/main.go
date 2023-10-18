package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PostProduct()
}

func PostProduct() {
	const url = "https://dummyjson.com/products/add"

	requestBody := strings.NewReader(`
		{
			"title": "Text "
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(nil)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
