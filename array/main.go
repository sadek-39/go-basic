package main

import "fmt"

func main() {
	fmt.Println("Array.....")

	var fruit [4]string

	fruit[0] = "Apple"
	fruit[3] = "Mango"
	fruit[2] = "Cake"

	fmt.Println("The fruit is : ", fruit)
	fmt.Println("The array length:", len(fruit))

	var vegList = [3]string{"alu", "tomato", "morich"}

	fmt.Println("The veg list is:", vegList)
}
