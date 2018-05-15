package main

import "fmt"

func main() {
	b := 2
	myAddOne := func(a int) {
		b = a + b
	}
	myAddOne(1)
	fmt.Println(b)
}
