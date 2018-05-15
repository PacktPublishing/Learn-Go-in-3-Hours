package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)
		i = i + i
		if i > 10 {
			break
		}
	}
	fmt.Println(i)
}
