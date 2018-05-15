package main

import "fmt"

func main() {
	s := "Hello, "
	var r rune = 127757
	s = s + string(r)
	fmt.Println(s)
}
