package main

import "fmt"

func main() {
	s := "Hello, world!"
	b := s[0]
	b2 := s[4]
	fmt.Println(s, b, string(b), b2, string(b2))
}
