package main

import "fmt"

func setTo10Fail(a *int) {
	a = new(int)
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)
}
