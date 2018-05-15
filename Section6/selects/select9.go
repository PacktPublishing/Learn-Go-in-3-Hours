package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workers := 10000
	pool := make(chan func(int) int, workers)
	for i := 0; i < workers; i++ {
		pool <- func(in int) int {
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < 100000; i++ {
		start := time.Now()
		select {
		case f := <-pool:
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)
		default:
			fmt.Println("rejecting request", i, "too busy")
		}
	}
	wg.Wait()
	fmt.Println("total processed:", count)
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
