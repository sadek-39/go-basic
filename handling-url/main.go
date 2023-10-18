package main

import (
	"fmt"
	"net/url"
)

const handleUrl = "https://dummyjson.com/products/search?q=phone&category=xyz"

func main() {
	fmt.Println("Handling url......")

	//parsing

	result, _ := url.Parse(handleUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query param %T\n", qparams)

	fmt.Println(qparams["q"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "dummyjson.com",
		Path:     "/products/search",
		RawQuery: "q=phone&category=xyz",
	}

	fmt.Println("The constructive url is : ", partsOfUrl.String())

}
