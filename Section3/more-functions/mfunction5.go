package main

import "fmt"

func main() {
	b := 2
	myAddOne := func(a int) int {
		return a + b
	}
	fmt.Println(myAddOne(1))
}
