package main

import "fmt"

func main() {
	s := "Hello, "
	r := '🌍'
	s = s + string(r)
	fmt.Println(s)
}
