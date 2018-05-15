package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	default:
		fmt.Println("nothing else works")
	}
}
