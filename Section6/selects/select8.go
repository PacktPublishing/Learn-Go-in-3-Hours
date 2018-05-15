package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	totalStart := time.Now()
	for i := 0; i < 100000; i++ {
		start := time.Now()
		wg.Add(1)
		go func(in int) {
			time.Sleep(1 * time.Second)
			out := 2 * in
			_ = out
			fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
