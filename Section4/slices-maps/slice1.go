package main

import "fmt"

func main() {
	s := make([]string, 0)
	fmt.Println("length of s:", len(s))
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	s2 := make([]string, 2)
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))
}
