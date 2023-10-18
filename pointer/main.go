package main

import "fmt"

func main() {

	number := 10

	ptr := &number

	fmt.Println("Value of pointer memory:", ptr)
	fmt.Println("Value of pointer value:", *ptr)

	*ptr = *ptr * 2

	fmt.Println("Value of pointer:", number)
	fmt.Println("Value of pointer value:", *ptr)
	fmt.Println("Value of pointer memory:", ptr)

	number = number + 2

	fmt.Println("Value of pointer:", number)
	fmt.Println("Value of pointer value:", *ptr)
	fmt.Println("Value of pointer memory:", ptr)

}
