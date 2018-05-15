package main

import "fmt"

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func main() {
	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(1))
	fmt.Println(addTwo(1))
}
