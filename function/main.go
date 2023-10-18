package main

import "fmt"

func main() {
	fmt.Println("Adding...")

	result := adder(2, 3)
	fmt.Println(result)
	fmt.Println("Pro Adder")
	fmt.Println(proAdder(7, 5, 6, 8))

}

func adder(x, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
