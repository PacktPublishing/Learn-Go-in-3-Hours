package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		b := a / 2
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is less than or equal to 5:")
	}
}
