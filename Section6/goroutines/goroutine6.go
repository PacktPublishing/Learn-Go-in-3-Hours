package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
