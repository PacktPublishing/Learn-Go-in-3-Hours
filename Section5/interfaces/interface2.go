package main

import "fmt"

func main() {
	var i interface{}
	i = "Hello"
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(j, k, ok)
	m := i.(int)
	fmt.Println(m)
}
