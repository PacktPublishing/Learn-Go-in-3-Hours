package main

import "fmt"

func main() {
	addNumbers(2, 3)
	addNumbers(4, 10)
	addNumbers(100, -100)
}

func addNumbers(a int, b int) {
	fmt.Println(a + b)
}

func addNumbers(a int, b int, c int) {
	fmt.Println(a + b + c)
}
