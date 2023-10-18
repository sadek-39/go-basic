package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

var upWebsites = []string{}

func main() {
	websites := []string{
		"https://google.com",
		"https://go.dev",
		"https://bkash.com",
		"https://djdj.com",
	}

	for _, val := range websites {
		go getStatus(val)
		wg.Add(1)
	}
	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatus(url string) {
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	} else {
		mut.Lock()
		upWebsites = append(upWebsites, url)
		mut.Unlock()
		fmt.Printf("%s has status code %d \n", url, res.StatusCode)
		fmt.Println("Website list which is up: ", upWebsites)
	}

}
