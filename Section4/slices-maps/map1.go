package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m:", h)
	fmt.Println("a in m:", m["a"])

	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m:", v)
	}

	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"])
}
