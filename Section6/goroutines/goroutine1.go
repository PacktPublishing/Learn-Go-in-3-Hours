package main

import "fmt"

func runMe() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	go runMe()
}
