package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channel")

	wg := &sync.WaitGroup{}
	mych := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)

		defer wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)

		defer wg.Done()
	}(mych, wg)

	wg.Wait()
}
