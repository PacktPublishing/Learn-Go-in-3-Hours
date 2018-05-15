package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	in <- 1
	o1 := <-out
	in <- 2
	o2 := <-out
	fmt.Println(o1, o2)
}
