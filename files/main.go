package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files....")
	file, err := os.Create("./test-file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, "Hello this is a testing file")

	if err != nil {
		panic(err)
	}

	fmt.Println("length is :", length)
	defer file.Close()
	readFile(file.Name())

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data from file \n", string(data))
}
