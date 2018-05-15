package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("So long!")
	} else if word == "greetings" {
		fmt.Println("Salutations")
	} else {
		fmt.Println("I don't know what you said")
	}
}
