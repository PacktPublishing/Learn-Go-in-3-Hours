package main

import (
	"fmt"
	"time"
)

func runMe() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	go runMe()
	time.Sleep(1 * time.Second)
}
