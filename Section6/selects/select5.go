package main

import "fmt"

func multiples(i int) chan int {
	out := make(chan int)
	curVal := 0
	go func() {
		for {
			out <- curVal * i
			curVal++
		}
	}()
	return out
}

func main() {
	twosCh := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}

	//do more stuff
}
