package main

import "fmt"

func main() {
	s3 := []string{"a", "b", "c"}
	s4 := s3[0:2]
	fmt.Println("s4:", s4)
	s3[0] = "d"
	fmt.Println("s4:", s4)

	var s5 []string
	s5 = s3
	s5[1] = "camel"
	fmt.Println("s3:", s3)

	modSlice(s3)
	fmt.Println("s3[0]:", s3[0])
}

func modSlice(s []string) {
	s[0] = "hello"
}
