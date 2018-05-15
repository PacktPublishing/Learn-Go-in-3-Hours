package main

import "fmt"

func main() {
	func addOne(a int) int {
		return a + 1
	}
	
	myAddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne(1))
}
