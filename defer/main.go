package main

import "fmt"

func main() {
	//defer fmt.Println("World")
	//defer fmt.Println("1")
	//defer fmt.Println("3")
	//fmt.Println("Hello")
	//fmt.Println("2")
	myDefer()

}

// Defer always follow last in  first out

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
