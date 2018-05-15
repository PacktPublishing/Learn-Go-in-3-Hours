package main

import "fmt"

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	addOne := makeAdder(1)
	doubleAddOne := makeDoubler(addOne)

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))
}
