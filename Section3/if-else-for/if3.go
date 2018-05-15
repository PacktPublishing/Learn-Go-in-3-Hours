package main

import "fmt"

func main() {
	a := 10
	if a = 5 {
		fmt.Println("a is bigger than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}
}
