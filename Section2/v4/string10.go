package main

import "fmt"

func main() {
	s := "👋 🌍"
	s2 := s[:1]
	s3 := s[2:]
	fmt.Println(s, len(s), s2, len(s2), s3, len(s3))
}
