package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/products"

func main() {
	fmt.Println("Handling request .....")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The data is :", string(data))
}
