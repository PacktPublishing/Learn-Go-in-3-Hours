package main

import (
	"fmt"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello to", name, "from a goroutine")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("Bob")

	wg.Wait()
}
