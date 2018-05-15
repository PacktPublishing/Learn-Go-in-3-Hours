package main

import "fmt"

func main() {
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	var vals2 [4]int = vals
	fmt.Println(vals, vals2)
}
