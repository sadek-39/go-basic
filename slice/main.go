package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruit = []string{"apple", "banana"}

	fmt.Printf("Type of fruit slice %T", fruit)
	fmt.Println()

	fruit = append(fruit, "Mango", "Jackfruit")

	fmt.Println("The fruit:", fruit[:3])

	highScore := make([]int, 4)

	highScore[0] = 100
	highScore[1] = 89
	highScore[3] = 76

	fmt.Println("The highscore list:", highScore)
	sort.Ints(highScore)
	fmt.Println("The highscore list:", highScore)
	fmt.Println("The highscore is sorted:", sort.IntsAreSorted(highScore))

	//TODO: Delete item from slice
}
