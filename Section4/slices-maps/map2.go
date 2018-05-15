package main

import "fmt"

func main() {
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("b in m2:", m2["b"])
	delete(m2, "b")
	fmt.Println("b in m2:", m2["b"])
}
