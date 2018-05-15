package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	var m3 map[string]int

	fmt.Println("goodbye in m:", m["goodbye"])
	m3 = m
	m3["goodbye"] = 400
	fmt.Println("goodbye in m3:", m3["goodbye"])
	fmt.Println("goodbye in m:", m["goodbye"])

	modMap(m)
	fmt.Println("cheese in m:", m["cheese"])
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
