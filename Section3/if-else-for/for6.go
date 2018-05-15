package main

import "fmt"

func main() {
	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
