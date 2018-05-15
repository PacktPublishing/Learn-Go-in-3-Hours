package main

import "fmt"

func main() {
	s := "Hello, world!"
	s2 := s[:5]
	s3 := s[7:]
	fmt.Println(s, len(s), s2, len(s2), s3, len(s3))
}
