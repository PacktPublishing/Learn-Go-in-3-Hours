package main

import "fmt"

func main() {
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	go func() {
		for {
			i, ok := <-in
			if !ok {
				close(out)
				break
			}
			out <- i * 2
		}
	}()

	for v := range out {
		fmt.Println(v)
	}
}
