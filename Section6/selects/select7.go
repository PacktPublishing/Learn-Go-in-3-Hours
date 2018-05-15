package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	in2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			in2 <- i
		}
		close(in2)
		wg.Done()
	}()

	go func() {
		count := 0
		for count < 2 {
			select {
			case i, ok := <-in:
				if !ok {
					count++
					in = nil
					continue
				}
				fmt.Println("from in, result is", i*2)
			case i, ok := <-in2:
				if !ok {
					count++
					in2 = nil
					continue
				}
				fmt.Println("from in2, result is", i+2)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
