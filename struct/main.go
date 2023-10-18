package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	sadek := User{"Sadek", "sadeksltn@gmail.com", true, 28}

	fmt.Println(sadek)
	fmt.Printf("The details %+v", sadek)
}
